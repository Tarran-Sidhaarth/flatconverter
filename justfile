# Justfile for converting proto files in cleaned/ to FlatBuffer schema files

# Directory containing cleaned proto files
cleaned_dir := "cleaned"

# Output directory for generated fbs files
fbs_output_dir := "flatbuffers"

# Convert all proto files to FlatBuffer schemas
all:
	@echo "🔄 Converting proto files in {{cleaned_dir}} to FlatBuffer schemas..."
	@mkdir -p {{fbs_output_dir}}
	#!/bin/bash
	success_count=0
	error_count=0
	skipped_count=0
	
	find {{cleaned_dir}} -name "*.proto" | while read proto_file; do
		# Get relative path to maintain directory structure
		rel_path="${proto_file#{{cleaned_dir}}/}"
		
		# Skip Google API files (like the Python script)
		if [[ "$rel_path" == *"google/api"* ]]; then
			echo "⊘ Skipping Google API file: $rel_path"
			continue
		fi
		
		# Create target directory structure
		target_dir="{{fbs_output_dir}}/$(dirname "$rel_path")"
		mkdir -p "$target_dir"
		
		# Generate FBS file using flatc
		echo "✓ Converting: $rel_path"
		if flatc --proto -I {{cleaned_dir}} -o "$target_dir" "$proto_file" 2>/dev/null; then
			# Post-process the generated .fbs file to fix includes
			fbs_file="$target_dir/$(basename "$proto_file" .proto).fbs"
			if [[ -f "$fbs_file" ]]; then
				# Convert import statements to include statements
				sed -i 's/import "\([^"]*\)\.proto"/include "\1.fbs"/g' "$fbs_file" 2>/dev/null || true
				echo "  ↳ Fixed includes in: $(basename "$fbs_file")"
			fi
		else
			echo "✗ Failed to convert: $rel_path"
		fi
	done
	@echo "📊 Conversion completed! Check {{fbs_output_dir}} for generated FlatBuffer schemas"

# Convert a specific proto file
convert FILE:
	@rel_path="{{FILE}}"
	@target_dir="{{fbs_output_dir}}/$(dirname "{{FILE}}")"
	@mkdir -p "$target_dir"
	@echo "Converting: {{FILE}}"
	@if flatc --proto -I {{cleaned_dir}} -o "$target_dir" "{{cleaned_dir}}/{{FILE}}"; then \
		fbs_file="$target_dir/$(basename "{{FILE}}" .proto).fbs"; \
		if [[ -f "$fbs_file" ]]; then \
			sed -i 's/import "\([^"]*\)\.proto"/include "\1.fbs"/g' "$fbs_file" 2>/dev/null || true; \
			echo "✓ Converted and fixed includes"; \
		fi; \
	else \
		echo "✗ Failed to convert {{FILE}}"; \
	fi

# List all proto files that will be converted (excluding Google API)
list:
	@echo "Proto files found in {{cleaned_dir}} (excluding Google API):"
	@find {{cleaned_dir}} -name '*.proto' | grep -v 'google/api' | sed 's|{{cleaned_dir}}/||'

# Clean generated fbs files
clean:
	@echo "🧹 Cleaning generated FlatBuffer files..."
	@rm -rf {{fbs_output_dir}}
	@echo "✓ Cleaned {{fbs_output_dir}}"

# Show statistics
stats:
	@echo "📊 Statistics:"
	@echo "  Total proto files: $(find {{cleaned_dir}} -name '*.proto' | wc -l)"
	@echo "  Google API files (will be skipped): $(find {{cleaned_dir}} -name '*.proto' | grep -c 'google/api' || echo 0)"
	@echo "  Files to convert: $(find {{cleaned_dir}} -name '*.proto' | grep -v 'google/api' | wc -l)"
	@if [[ -d "{{fbs_output_dir}}" ]]; then \
		echo "  Generated FBS files: $(find {{fbs_output_dir}} -name '*.fbs' | wc -l)"; \
	fi

# Help
help:
	@echo "Available commands:"
	@echo "  just all           # Convert all proto files to FlatBuffer schemas"
	@echo "  just convert FILE  # Convert a specific proto file (relative path)"
	@echo "  just list          # List all proto files (excluding Google API)"
	@echo "  just stats         # Show conversion statistics"
	@echo "  just clean         # Remove generated FlatBuffer schema files"
	@echo "  just help          # Show this help message"

# Set default recipe
default: all
