# DO NOT EDIT!
# python generated by Buffman 💪
# Versions:
# 		Buffman: 1.0.0
# 		Flatc: v25.2.10

# namespace: services

import flatbuffers
from flatbuffers.compat import import_numpy
np = import_numpy()

class UserAnalyticsRequest(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAs(cls, buf, offset=0):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = UserAnalyticsRequest()
        x.Init(buf, n + offset)
        return x

    @classmethod
    def GetRootAsUserAnalyticsRequest(cls, buf, offset=0):
        """This method is deprecated. Please switch to GetRootAs."""
        return cls.GetRootAs(buf, offset)
    # UserAnalyticsRequest
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # UserAnalyticsRequest
    def UserId(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            return self._tab.String(o + self._tab.Pos)
        return None

    # UserAnalyticsRequest
    def StartDate(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(6))
        if o != 0:
            x = self._tab.Indirect(o + self._tab.Pos)
            from common.Timestamp import Timestamp
            obj = Timestamp()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

    # UserAnalyticsRequest
    def EndDate(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(8))
        if o != 0:
            x = self._tab.Indirect(o + self._tab.Pos)
            from common.Timestamp import Timestamp
            obj = Timestamp()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

def UserAnalyticsRequestStart(builder):
    builder.StartObject(3)

def Start(builder):
    UserAnalyticsRequestStart(builder)

def UserAnalyticsRequestAddUserId(builder, userId):
    builder.PrependUOffsetTRelativeSlot(0, flatbuffers.number_types.UOffsetTFlags.py_type(userId), 0)

def AddUserId(builder, userId):
    UserAnalyticsRequestAddUserId(builder, userId)

def UserAnalyticsRequestAddStartDate(builder, startDate):
    builder.PrependUOffsetTRelativeSlot(1, flatbuffers.number_types.UOffsetTFlags.py_type(startDate), 0)

def AddStartDate(builder, startDate):
    UserAnalyticsRequestAddStartDate(builder, startDate)

def UserAnalyticsRequestAddEndDate(builder, endDate):
    builder.PrependUOffsetTRelativeSlot(2, flatbuffers.number_types.UOffsetTFlags.py_type(endDate), 0)

def AddEndDate(builder, endDate):
    UserAnalyticsRequestAddEndDate(builder, endDate)

def UserAnalyticsRequestEnd(builder):
    return builder.EndObject()

def End(builder):
    return UserAnalyticsRequestEnd(builder)
