# DO NOT EDIT!
# python generated by Buffman 💪
# Versions:
# 		Buffman: 1.0.0
# 		Flatc: v25.2.10

# namespace: services

import flatbuffers
from flatbuffers.compat import import_numpy
np = import_numpy()

class CreateUserResponse(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAs(cls, buf, offset=0):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = CreateUserResponse()
        x.Init(buf, n + offset)
        return x

    @classmethod
    def GetRootAsCreateUserResponse(cls, buf, offset=0):
        """This method is deprecated. Please switch to GetRootAs."""
        return cls.GetRootAs(buf, offset)
    # CreateUserResponse
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # CreateUserResponse
    def User(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            x = self._tab.Indirect(o + self._tab.Pos)
            from services.User import User
            obj = User()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

    # CreateUserResponse
    def Success(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(6))
        if o != 0:
            return bool(self._tab.Get(flatbuffers.number_types.BoolFlags, o + self._tab.Pos))
        return False

    # CreateUserResponse
    def Message(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(8))
        if o != 0:
            return self._tab.String(o + self._tab.Pos)
        return None

def CreateUserResponseStart(builder):
    builder.StartObject(3)

def Start(builder):
    CreateUserResponseStart(builder)

def CreateUserResponseAddUser(builder, user):
    builder.PrependUOffsetTRelativeSlot(0, flatbuffers.number_types.UOffsetTFlags.py_type(user), 0)

def AddUser(builder, user):
    CreateUserResponseAddUser(builder, user)

def CreateUserResponseAddSuccess(builder, success):
    builder.PrependBoolSlot(1, success, 0)

def AddSuccess(builder, success):
    CreateUserResponseAddSuccess(builder, success)

def CreateUserResponseAddMessage(builder, message):
    builder.PrependUOffsetTRelativeSlot(2, flatbuffers.number_types.UOffsetTFlags.py_type(message), 0)

def AddMessage(builder, message):
    CreateUserResponseAddMessage(builder, message)

def CreateUserResponseEnd(builder):
    return builder.EndObject()

def End(builder):
    return CreateUserResponseEnd(builder)
