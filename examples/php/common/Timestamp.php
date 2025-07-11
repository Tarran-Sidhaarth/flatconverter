// DO NOT EDIT!
// php generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10
// automatically generated by the FlatBuffers compiler, do not modify

namespace common;

use \Google\FlatBuffers\Struct;
use \Google\FlatBuffers\Table;
use \Google\FlatBuffers\ByteBuffer;
use \Google\FlatBuffers\FlatBufferBuilder;

class Timestamp extends Table
{
    /**
     * @param ByteBuffer $bb
     * @return Timestamp
     */
    public static function getRootAsTimestamp(ByteBuffer $bb)
    {
        $obj = new Timestamp();
        return ($obj->init($bb->getInt($bb->getPosition()) + $bb->getPosition(), $bb));
    }

    /**
     * @param int $_i offset
     * @param ByteBuffer $_bb
     * @return Timestamp
     **/
    public function init($_i, ByteBuffer $_bb)
    {
        $this->bb_pos = $_i;
        $this->bb = $_bb;
        return $this;
    }

    /**
     * @return long
     */
    public function getSeconds()
    {
        $o = $this->__offset(4);
        return $o != 0 ? $this->bb->getLong($o + $this->bb_pos) : 0;
    }

    /**
     * @return int
     */
    public function getNanos()
    {
        $o = $this->__offset(6);
        return $o != 0 ? $this->bb->getInt($o + $this->bb_pos) : 0;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return void
     */
    public static function startTimestamp(FlatBufferBuilder $builder)
    {
        $builder->StartObject(2);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return Timestamp
     */
    public static function createTimestamp(FlatBufferBuilder $builder, $seconds, $nanos)
    {
        $builder->startObject(2);
        self::addSeconds($builder, $seconds);
        self::addNanos($builder, $nanos);
        $o = $builder->endObject();
        return $o;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param long
     * @return void
     */
    public static function addSeconds(FlatBufferBuilder $builder, $seconds)
    {
        $builder->addLongX(0, $seconds, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param int
     * @return void
     */
    public static function addNanos(FlatBufferBuilder $builder, $nanos)
    {
        $builder->addIntX(1, $nanos, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return int table offset
     */
    public static function endTimestamp(FlatBufferBuilder $builder)
    {
        $o = $builder->endObject();
        return $o;
    }
}
