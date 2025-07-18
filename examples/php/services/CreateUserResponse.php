// DO NOT EDIT!
// php generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10
// automatically generated by the FlatBuffers compiler, do not modify

namespace services;

use \Google\FlatBuffers\Struct;
use \Google\FlatBuffers\Table;
use \Google\FlatBuffers\ByteBuffer;
use \Google\FlatBuffers\FlatBufferBuilder;

class CreateUserResponse extends Table
{
    /**
     * @param ByteBuffer $bb
     * @return CreateUserResponse
     */
    public static function getRootAsCreateUserResponse(ByteBuffer $bb)
    {
        $obj = new CreateUserResponse();
        return ($obj->init($bb->getInt($bb->getPosition()) + $bb->getPosition(), $bb));
    }

    /**
     * @param int $_i offset
     * @param ByteBuffer $_bb
     * @return CreateUserResponse
     **/
    public function init($_i, ByteBuffer $_bb)
    {
        $this->bb_pos = $_i;
        $this->bb = $_bb;
        return $this;
    }

    public function getUser()
    {
        $obj = new User();
        $o = $this->__offset(4);
        return $o != 0 ? $obj->init($this->__indirect($o + $this->bb_pos), $this->bb) : 0;
    }

    /**
     * @return bool
     */
    public function getSuccess()
    {
        $o = $this->__offset(6);
        return $o != 0 ? $this->bb->getBool($o + $this->bb_pos) : false;
    }

    public function getMessage()
    {
        $o = $this->__offset(8);
        return $o != 0 ? $this->__string($o + $this->bb_pos) : null;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return void
     */
    public static function startCreateUserResponse(FlatBufferBuilder $builder)
    {
        $builder->StartObject(3);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return CreateUserResponse
     */
    public static function createCreateUserResponse(FlatBufferBuilder $builder, $user, $success, $message)
    {
        $builder->startObject(3);
        self::addUser($builder, $user);
        self::addSuccess($builder, $success);
        self::addMessage($builder, $message);
        $o = $builder->endObject();
        return $o;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param VectorOffset
     * @return void
     */
    public static function addUser(FlatBufferBuilder $builder, $user)
    {
        $builder->addOffsetX(0, $user, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param bool
     * @return void
     */
    public static function addSuccess(FlatBufferBuilder $builder, $success)
    {
        $builder->addBoolX(1, $success, false);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param StringOffset
     * @return void
     */
    public static function addMessage(FlatBufferBuilder $builder, $message)
    {
        $builder->addOffsetX(2, $message, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return int table offset
     */
    public static function endCreateUserResponse(FlatBufferBuilder $builder)
    {
        $o = $builder->endObject();
        return $o;
    }
}
