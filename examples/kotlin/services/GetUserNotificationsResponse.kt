// DO NOT EDIT!
// kotlin generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10

package services

import com.google.flatbuffers.BaseVector
import com.google.flatbuffers.BooleanVector
import com.google.flatbuffers.ByteVector
import com.google.flatbuffers.Constants
import com.google.flatbuffers.DoubleVector
import com.google.flatbuffers.FlatBufferBuilder
import com.google.flatbuffers.FloatVector
import com.google.flatbuffers.LongVector
import com.google.flatbuffers.StringVector
import com.google.flatbuffers.Struct
import com.google.flatbuffers.Table
import com.google.flatbuffers.UnionVector
import java.nio.ByteBuffer
import java.nio.ByteOrder
import kotlin.math.sign

@Suppress("unused")
class GetUserNotificationsResponse : Table() {

    fun __init(_i: Int, _bb: ByteBuffer)  {
        __reset(_i, _bb)
    }
    fun __assign(_i: Int, _bb: ByteBuffer) : GetUserNotificationsResponse {
        __init(_i, _bb)
        return this
    }
    fun notifications(j: Int) : services.Notification? = notifications(services.Notification(), j)
    fun notifications(obj: services.Notification, j: Int) : services.Notification? {
        val o = __offset(4)
        return if (o != 0) {
            obj.__assign(__indirect(__vector(o) + j * 4), bb)
        } else {
            null
        }
    }
    val notificationsLength : Int
        get() {
            val o = __offset(4); return if (o != 0) __vector_len(o) else 0
        }
    val totalCount : Int
        get() {
            val o = __offset(6)
            return if(o != 0) bb.getInt(o + bb_pos) else 0
        }
    val userInfo : services.User? get() = userInfo(services.User())
    fun userInfo(obj: services.User) : services.User? {
        val o = __offset(8)
        return if (o != 0) {
            obj.__assign(__indirect(o + bb_pos), bb)
        } else {
            null
        }
    }
    companion object {
        fun validateVersion() = Constants.FLATBUFFERS_25_2_10()
        fun getRootAsGetUserNotificationsResponse(_bb: ByteBuffer): GetUserNotificationsResponse = getRootAsGetUserNotificationsResponse(_bb, GetUserNotificationsResponse())
        fun getRootAsGetUserNotificationsResponse(_bb: ByteBuffer, obj: GetUserNotificationsResponse): GetUserNotificationsResponse {
            _bb.order(ByteOrder.LITTLE_ENDIAN)
            return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb))
        }
        fun createGetUserNotificationsResponse(builder: FlatBufferBuilder, notificationsOffset: Int, totalCount: Int, userInfoOffset: Int) : Int {
            builder.startTable(3)
            addUserInfo(builder, userInfoOffset)
            addTotalCount(builder, totalCount)
            addNotifications(builder, notificationsOffset)
            return endGetUserNotificationsResponse(builder)
        }
        fun startGetUserNotificationsResponse(builder: FlatBufferBuilder) = builder.startTable(3)
        fun addNotifications(builder: FlatBufferBuilder, notifications: Int) = builder.addOffset(0, notifications, 0)
        fun createNotificationsVector(builder: FlatBufferBuilder, data: IntArray) : Int {
            builder.startVector(4, data.size, 4)
            for (i in data.size - 1 downTo 0) {
                builder.addOffset(data[i])
            }
            return builder.endVector()
        }
        fun startNotificationsVector(builder: FlatBufferBuilder, numElems: Int) = builder.startVector(4, numElems, 4)
        fun addTotalCount(builder: FlatBufferBuilder, totalCount: Int) = builder.addInt(1, totalCount, 0)
        fun addUserInfo(builder: FlatBufferBuilder, userInfo: Int) = builder.addOffset(2, userInfo, 0)
        fun endGetUserNotificationsResponse(builder: FlatBufferBuilder) : Int {
            val o = builder.endTable()
            return o
        }
    }
}
