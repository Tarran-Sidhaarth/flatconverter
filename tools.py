import os
import re
import glob
import shutil
import subprocess
import argparse
import tempfile
from pathlib import Path

class EnhancedProtoApiCleaner:
    def __init__(self, input_dir, output_dir, keep_cleaned=False):
        self.input_dir = Path(input_dir).resolve()
        self.output_dir = Path(output_dir).resolve()
        self.keep_cleaned = keep_cleaned
        
        # Folders to ignore during FlatBuffers generation
        self.ignore_folders = {
            'google/api',
            'google/cloud',
            'google/logging',
            'google/monitoring',
            'google/protobuf/descriptor'  # Exclude descriptor files
        }
        
        # Safety check - ensure we never modify original files
        if self.output_dir == self.input_dir:
            raise ValueError("Output directory cannot be the same as input directory!")
        
        # Create output directory
        self.output_dir.mkdir(parents=True, exist_ok=True)
    
    def should_ignore_for_fbs_generation(self, file_path):
        """Check if file should be ignored during FlatBuffers generation"""
        file_str = str(file_path)
        
        # Check if file is in any of the ignore folders
        for ignore_folder in self.ignore_folders:
            if ignore_folder in file_str:
                return True
        
        # Always process google/protobuf files (except descriptor)
        if 'google/protobuf' in file_str and 'descriptor' not in file_str:
            return False
            
        # Always process google/longrunning files (we clean these specially)
        if 'google/longrunning' in file_str:
            return False
            
        return False
    
    def is_longrunning_operations_file(self, file_path):
        """Check if this is a google/longrunning/operations.proto file"""
        return "google/longrunning/operations.proto" in str(file_path) or file_path.name == "operations.proto"
    
    def clean_longrunning_operations(self, content):
        """Special cleaning for google/longrunning/operations.proto - keep only Operation message"""
        
        # Extract the Operation message and its dependencies
        operation_message_pattern = r'(// This resource represents a long-running operation.*?message Operation \{.*?\n\})'
        operation_match = re.search(operation_message_pattern, content, re.DOTALL)
        
        if not operation_match:
            # Fallback pattern if the comment is different
            operation_message_pattern = r'(message Operation \{.*?\n\})'
            operation_match = re.search(operation_message_pattern, content, re.DOTALL)
        
        if operation_match:
            # Keep only essential parts
            cleaned_content = []
            
            # Keep syntax and package
            syntax_match = re.search(r'syntax = "proto3";', content)
            if syntax_match:
                cleaned_content.append(syntax_match.group(0))
            
            package_match = re.search(r'package [^;]+;', content)
            if package_match:
                cleaned_content.append(package_match.group(0))
            
            # Keep only necessary imports for Operation message (excluding descriptor)
            necessary_imports = [
                'import "google/protobuf/any.proto";',
                'import "google/rpc/status.proto";'
            ]
            
            for imp in necessary_imports:
                if imp in content:
                    cleaned_content.append(imp)
            
            # Add the Operation message
            cleaned_content.append(operation_match.group(0))
            
            return '\n\n'.join(cleaned_content) + '\n'
        
        # If Operation message not found, return original content with standard cleaning
        return self.clean_proto_content(content)
    
    def clean_proto_content(self, content):
        """Enhanced cleaning of Google API annotations and options from proto content"""
        
        # Remove Google API imports
        content = re.sub(r'import\s+"google/api/annotations\.proto";\s*\n', '', content)
        content = re.sub(r'import\s+"google/api/field_behavior\.proto";\s*\n', '', content)
        content = re.sub(r'import\s+"google/api/client\.proto";\s*\n', '', content)
        content = re.sub(r'import\s+"google/api/resource\.proto";\s*\n', '', content)
        content = re.sub(r'import\s+"google/api/http\.proto";\s*\n', '', content)
        content = re.sub(r'import\s+"google/api/field_info\.proto";\s*\n', '', content)
        
        # Remove google/protobuf/descriptor.proto import - ENHANCED
        content = re.sub(r'import\s+"google/protobuf/descriptor\.proto";\s*\n', '', content)
        
        # Remove any descriptor-related extensions or options
        content = re.sub(r'extend\s+google\.protobuf\..*?;', '', content, flags=re.MULTILINE)
        content = re.sub(r'option\s*\([^)]*descriptor[^)]*\)[^;]*;', '', content, flags=re.MULTILINE)
        
        # Remove HTTP options from RPC methods
        content = re.sub(
            r'option\s*\(google\.api\.http\)\s*=\s*\{[^{}]*(?:\{[^{}]*\}[^{}]*)*\};\s*\n',
            '',
            content,
            flags=re.MULTILINE | re.DOTALL
        )
        
        # Remove field behavior annotations
        content = re.sub(
            r'\[\s*\(google\.api\.field_behavior\)\s*=\s*[^\]]*\]',
            '',
            content
        )
        
        # Remove resource_reference annotations - ENHANCED
        content = re.sub(
            r'\[\s*\(google\.api\.resource_reference\)\s*[^]]*\]',
            '',
            content,
            flags=re.MULTILINE | re.DOTALL
        )
        
        # Remove complex resource_reference with .type syntax
        content = re.sub(
            r'\[\s*\(google\.api\.resource_reference\)\.type\s*=\s*[^\]]*\]',
            '',
            content
        )
        
        # Remove message-level resource options
        content = re.sub(
            r'option\s*\(google\.api\.resource\)\s*=\s*\{[^{}]*(?:\{[^{}]*\}[^{}]*)*\};\s*\n',
            '',
            content,
            flags=re.MULTILINE | re.DOTALL
        )
        
        # Remove method signature options
        content = re.sub(
            r'option\s*\(google\.api\.method_signature\)\s*=\s*"[^"]*";\s*\n',
            '',
            content
        )
        
        # Remove oauth scopes options
        content = re.sub(
            r'option\s*\(google\.api\.oauth_scopes\)\s*=\s*"[^"]*";\s*\n',
            '',
            content
        )
        
        # Remove default_host options
        content = re.sub(
            r'option\s*\(google\.api\.default_host\)\s*=\s*"[^"]*";\s*\n',
            '',
            content
        )
        
        # Clean up RPC method definitions - remove empty option blocks
        content = re.sub(
            r'(\s*rpc\s+\w+\s*\([^)]*\)\s*returns\s*\([^)]*\)\s*)\{\s*\}',
            r'\1{}',
            content
        )
        
        # Clean up extra whitespace and empty lines
        content = re.sub(r'\n\s*\n\s*\n', '\n\n', content)
        content = re.sub(r'[ \t]+\n', '\n', content)  # Remove trailing spaces
        
        return content
    
    def convert_imports_to_includes(self, fbs_content):
        """Convert proto imports to FlatBuffers includes"""
        # Pattern to match: import "folder1/file.proto";
        # Replace with: include "folder1/file.fbs";
        fbs_content = re.sub(
            r'import\s+"([^"]+)\.proto"\s*;',
            r'include "\1.fbs";',
            fbs_content
        )
        return fbs_content
    
    def process_proto_file(self, input_file, output_file):
        """Process a single proto file"""
        try:
            # Read original file (read-only)
            with open(input_file, 'r', encoding='utf-8') as f:
                original_content = f.read()
            
            # Check if this is a longrunning operations file
            if self.is_longrunning_operations_file(input_file):
                print(f"🔧 Special handling for longrunning operations file: {input_file.name}")
                cleaned_content = self.clean_longrunning_operations(original_content)
            else:
                # Standard cleaning
                cleaned_content = self.clean_proto_content(original_content)
            
            # Create output directory if needed
            output_file.parent.mkdir(parents=True, exist_ok=True)
            
            # Write cleaned content to output file
            with open(output_file, 'w', encoding='utf-8') as f:
                f.write(cleaned_content)
            
            # Log the changes
            if cleaned_content != original_content:
                print(f"✓ Cleaned: {input_file.name} -> {output_file}")
            else:
                print(f"○ No changes needed: {input_file.name}")
                
            return True
            
        except Exception as e:
            print(f"✗ Error processing {input_file}: {e}")
            return False
    
    def process_directory(self):
        """Process all proto files in the input directory"""
        proto_files = list(self.input_dir.rglob("*.proto"))
        
        if not proto_files:
            print(f"⚠️  No .proto files found in {self.input_dir}")
            return 0
        
        processed_count = 0
        
        for proto_file in proto_files:
            # Skip descriptor files entirely
            if 'descriptor' in str(proto_file):
                print(f"⊘ Skipping descriptor file: {proto_file.name}")
                continue
                
            # Maintain directory structure in output
            rel_path = proto_file.relative_to(self.input_dir)
            output_file = self.output_dir / rel_path
            
            if self.process_proto_file(proto_file, output_file):
                processed_count += 1
        
        print(f"\n📊 Processing Summary:")
        print(f"   Total files found: {len(proto_files)}")
        print(f"   Successfully processed: {processed_count}")
        print(f"   Output directory: {self.output_dir}")
        
        return processed_count
    
    def generate_flatbuffers_with_structure(self, fbs_output_dir):
        """Generate FlatBuffers files with proper include paths and directory structure"""
        fbs_output_path = Path(fbs_output_dir)
        fbs_output_path.mkdir(parents=True, exist_ok=True)
        
        proto_files = list(self.output_dir.rglob("*.proto"))
        success_count = 0
        error_count = 0
        skipped_count = 0
        
        print(f"\n🔄 Generating FlatBuffers schemas (ignoring Google API folders)...")
        
        for proto_file in proto_files:
            try:
                # Check if this file should be ignored for FBS generation
                if self.should_ignore_for_fbs_generation(proto_file):
                    rel_path = proto_file.relative_to(self.output_dir)
                    print(f"⊘ Skipping Google API file: {rel_path}")
                    skipped_count += 1
                    continue
                
                # Get relative path to maintain directory structure
                rel_path = proto_file.relative_to(self.output_dir)
                
                # Create the target directory structure
                target_dir = fbs_output_path / rel_path.parent
                target_dir.mkdir(parents=True, exist_ok=True)
                
                # Generate FBS file using flatc with include path
                cmd = f"flatc --proto -I {self.output_dir} -o {target_dir} {proto_file}"
                result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
                
                if result.returncode == 0:
                    # Find and post-process the generated .fbs file
                    fbs_filename = proto_file.stem + ".fbs"
                    generated_fbs = target_dir / fbs_filename
                    
                    if generated_fbs.exists():
                        self.fix_fbs_includes(generated_fbs)
                        print(f"✓ Generated FBS: {rel_path}")
                        success_count += 1
                    else:
                        print(f"✗ FBS file not found after generation: {rel_path}")
                        error_count += 1
                else:
                    print(f"✗ Failed to generate FBS: {rel_path}")
                    if result.stderr:
                        print(f"   Error: {result.stderr.strip()}")
                    error_count += 1
                    
            except Exception as e:
                print(f"✗ Error processing {proto_file.name}: {e}")
                error_count += 1
        
        print(f"\n📊 FlatBuffers Generation Summary:")
        print(f"   Successful: {success_count} files")
        print(f"   Skipped (Google API): {skipped_count} files")
        print(f"   Failed: {error_count} files")
        print(f"   FBS output: {fbs_output_path}")
        
        return success_count > 0
    
    def fix_fbs_includes(self, fbs_file):
        """Fix includes in generated FBS files"""
        try:
            with open(fbs_file, 'r', encoding='utf-8') as f:
                content = f.read()
            
            # Convert any remaining import statements to include statements
            # and change .proto extensions to .fbs
            original_content = content
            content = self.convert_imports_to_includes(content)
            
            # Also handle include statements that might have been generated
            content = re.sub(
                r'include\s+"([^"]+)\.proto"',
                r'include "\1.fbs"',
                content
            )
            
            # Write back if changes were made
            if content != original_content:
                with open(fbs_file, 'w', encoding='utf-8') as f:
                    f.write(content)
                print(f"  ↳ Fixed includes in: {fbs_file.name}")
                
        except Exception as e:
            print(f"  ↳ Error fixing includes in {fbs_file}: {e}")
    
    def cleanup_directories(self, deps_dir):
        """Clean up temporary directories based on configuration"""
        # Clean up nogoogleapi directory (we don't need it after FBS generation)
        if not self.keep_cleaned and self.output_dir.exists():
            shutil.rmtree(self.output_dir)
            print(f"🧹 Cleaned up cleaned proto directory: {self.output_dir}")

def main():
    import sys
    keep_dependency = False
    for arg in sys.argv[1:]:
        if arg.startswith('--keep_dependency='):
            val = arg.split('=', 1)[1].lower()
            keep_dependency = val in ('1', 'true', 'yes')
    
    """Main function with hardcoded directories"""
    
    # HARDCODED DIRECTORIES
    deps_dir = "/workspaces/protoverse/deps"
    cleaned_dir = "/workspaces/protoverse/nogoogleapi"
    fbs_dir = "/workspaces/protoverse/flatbuffers"
    
    print("🚀 Enhanced Proto File Cleaner with Hardcoded Directories")
    print("=" * 60)
    print(f"📁 Dependencies: {deps_dir}")
    print(f"📁 Cleaned Proto: {cleaned_dir}")
    print(f"📁 FlatBuffers: {fbs_dir}")
    print(f"🗑️  Delete deps after: {'No' if keep_dependency else 'Yes'}")
    print("=" * 60)
    
    try:
        # Step 1: Run buf export
        deps_path = Path(deps_dir)
        if deps_path.exists():
            shutil.rmtree(deps_path)
        
        print("🔄 Running buf export to get dependencies...")
        try:
            cmd = f"buf export . --output {deps_dir} --exclude-imports=false"
            result = subprocess.run(cmd, shell=True, capture_output=True, text=True, cwd=".")
            
            if result.returncode == 0:
                print(f"✓ Successfully exported dependencies to {deps_dir}")
            else:
                print(f"✗ Failed to run buf export")
                print(f"   Error: {result.stderr}")
                return 1
                
        except Exception as e:
            print(f"✗ Error running buf export: {e}")
            return 1
        
        # Verify deps directory exists
        if not Path(deps_dir).exists():
            print(f"❌ Dependencies directory does not exist: {deps_dir}")
            return 1
        
        # Step 2: Initialize cleaner with hardcoded directories
        cleaner = EnhancedProtoApiCleaner(deps_dir, cleaned_dir, keep_cleaned=False)
        
        # Process proto files
        processed_count = cleaner.process_directory()
        
        if processed_count > 0:
            # Step 3: Generate FlatBuffers files
            success = cleaner.generate_flatbuffers_with_structure(fbs_dir)
            
            if success:
                print(f"\n✅ Process completed successfully!")
                print(f"   Dependencies: {deps_dir} ({'kept' if keep_dependency else 'will be cleaned up'})")
                print(f"   Cleaned files: {cleaned_dir} (will be cleaned up)")
                print(f"   FlatBuffers: {fbs_dir}")
                
                # Step 4: Cleanup directories
                cleaner.cleanup_directories(deps_dir)
                if not keep_dependency and Path(deps_dir).exists():
                    shutil.rmtree(deps_dir)
                    print(f"🧹 Cleaned up dependencies directory: {deps_dir}")
                else:
                    print(f"📦 Dependencies directory kept: {deps_dir}")
                
                return 0
            else:
                print("❌ FlatBuffers generation failed")
                return 1
        else:
            print("⚠️  No files were processed.")
            return 1
            
    except KeyboardInterrupt:
        print("\n⚠️  Process interrupted by user")
        return 1
    except Exception as e:
        print(f"❌ Error: {e}")
        return 1

if __name__ == "__main__":
    exit(main())
