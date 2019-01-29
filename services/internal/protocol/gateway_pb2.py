# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gateway.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='gateway.proto',
  package='protocol',
  syntax='proto3',
  serialized_pb=_b('\n\rgateway.proto\x12\x08protocol\"4\n\x0b\x43MD_GATEWAY\"%\n\x04\x45NUM\x12\x0b\n\x07INVALID\x10\x00\x12\x10\n\x0cVERIFY_TOKEN\x10\x01\"T\n\x1f\x45NUM_GATEWAY_VERIFY_TOKEN_ERROR\"1\n\x04\x45NUM\x12\x06\n\x02OK\x10\x00\x12\x0f\n\x0bVERIFY_FAIL\x10\x01\x12\x10\n\x0cSYSTEM_ERROR\x10\x02\"2\n\x0eROLE_BASE_INFO\x12\x0e\n\x06RoleID\x18\x01 \x01(\x04\x12\x10\n\x08Userdata\x18\x02 \x01(\x0c\":\n\x18MSG_GATEWAY_VERIFY_TOKEN\x12\x0f\n\x07\x41\x63\x63ount\x18\x01 \x01(\t\x12\r\n\x05Token\x18\x02 \x01(\t\"\x82\x01\n\x1fMSG_GATEWAY_VERIFY_TOKEN_RESULT\x12\x36\n\x03\x45rr\x18\x01 \x01(\x0b\x32).protocol.ENUM_GATEWAY_VERIFY_TOKEN_ERROR\x12\'\n\x05Roles\x18\x02 \x03(\x0b\x32\x18.protocol.ROLE_BASE_INFOb\x06proto3')
)



_CMD_GATEWAY_ENUM = _descriptor.EnumDescriptor(
  name='ENUM',
  full_name='protocol.CMD_GATEWAY.ENUM',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='INVALID', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VERIFY_TOKEN', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=42,
  serialized_end=79,
)
_sym_db.RegisterEnumDescriptor(_CMD_GATEWAY_ENUM)

_ENUM_GATEWAY_VERIFY_TOKEN_ERROR_ENUM = _descriptor.EnumDescriptor(
  name='ENUM',
  full_name='protocol.ENUM_GATEWAY_VERIFY_TOKEN_ERROR.ENUM',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VERIFY_FAIL', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SYSTEM_ERROR', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=116,
  serialized_end=165,
)
_sym_db.RegisterEnumDescriptor(_ENUM_GATEWAY_VERIFY_TOKEN_ERROR_ENUM)


_CMD_GATEWAY = _descriptor.Descriptor(
  name='CMD_GATEWAY',
  full_name='protocol.CMD_GATEWAY',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _CMD_GATEWAY_ENUM,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=79,
)


_ENUM_GATEWAY_VERIFY_TOKEN_ERROR = _descriptor.Descriptor(
  name='ENUM_GATEWAY_VERIFY_TOKEN_ERROR',
  full_name='protocol.ENUM_GATEWAY_VERIFY_TOKEN_ERROR',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _ENUM_GATEWAY_VERIFY_TOKEN_ERROR_ENUM,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=81,
  serialized_end=165,
)


_ROLE_BASE_INFO = _descriptor.Descriptor(
  name='ROLE_BASE_INFO',
  full_name='protocol.ROLE_BASE_INFO',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='RoleID', full_name='protocol.ROLE_BASE_INFO.RoleID', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Userdata', full_name='protocol.ROLE_BASE_INFO.Userdata', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=167,
  serialized_end=217,
)


_MSG_GATEWAY_VERIFY_TOKEN = _descriptor.Descriptor(
  name='MSG_GATEWAY_VERIFY_TOKEN',
  full_name='protocol.MSG_GATEWAY_VERIFY_TOKEN',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Account', full_name='protocol.MSG_GATEWAY_VERIFY_TOKEN.Account', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Token', full_name='protocol.MSG_GATEWAY_VERIFY_TOKEN.Token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=219,
  serialized_end=277,
)


_MSG_GATEWAY_VERIFY_TOKEN_RESULT = _descriptor.Descriptor(
  name='MSG_GATEWAY_VERIFY_TOKEN_RESULT',
  full_name='protocol.MSG_GATEWAY_VERIFY_TOKEN_RESULT',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Err', full_name='protocol.MSG_GATEWAY_VERIFY_TOKEN_RESULT.Err', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Roles', full_name='protocol.MSG_GATEWAY_VERIFY_TOKEN_RESULT.Roles', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=280,
  serialized_end=410,
)

_CMD_GATEWAY_ENUM.containing_type = _CMD_GATEWAY
_ENUM_GATEWAY_VERIFY_TOKEN_ERROR_ENUM.containing_type = _ENUM_GATEWAY_VERIFY_TOKEN_ERROR
_MSG_GATEWAY_VERIFY_TOKEN_RESULT.fields_by_name['Err'].message_type = _ENUM_GATEWAY_VERIFY_TOKEN_ERROR
_MSG_GATEWAY_VERIFY_TOKEN_RESULT.fields_by_name['Roles'].message_type = _ROLE_BASE_INFO
DESCRIPTOR.message_types_by_name['CMD_GATEWAY'] = _CMD_GATEWAY
DESCRIPTOR.message_types_by_name['ENUM_GATEWAY_VERIFY_TOKEN_ERROR'] = _ENUM_GATEWAY_VERIFY_TOKEN_ERROR
DESCRIPTOR.message_types_by_name['ROLE_BASE_INFO'] = _ROLE_BASE_INFO
DESCRIPTOR.message_types_by_name['MSG_GATEWAY_VERIFY_TOKEN'] = _MSG_GATEWAY_VERIFY_TOKEN
DESCRIPTOR.message_types_by_name['MSG_GATEWAY_VERIFY_TOKEN_RESULT'] = _MSG_GATEWAY_VERIFY_TOKEN_RESULT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CMD_GATEWAY = _reflection.GeneratedProtocolMessageType('CMD_GATEWAY', (_message.Message,), dict(
  DESCRIPTOR = _CMD_GATEWAY,
  __module__ = 'gateway_pb2'
  # @@protoc_insertion_point(class_scope:protocol.CMD_GATEWAY)
  ))
_sym_db.RegisterMessage(CMD_GATEWAY)

ENUM_GATEWAY_VERIFY_TOKEN_ERROR = _reflection.GeneratedProtocolMessageType('ENUM_GATEWAY_VERIFY_TOKEN_ERROR', (_message.Message,), dict(
  DESCRIPTOR = _ENUM_GATEWAY_VERIFY_TOKEN_ERROR,
  __module__ = 'gateway_pb2'
  # @@protoc_insertion_point(class_scope:protocol.ENUM_GATEWAY_VERIFY_TOKEN_ERROR)
  ))
_sym_db.RegisterMessage(ENUM_GATEWAY_VERIFY_TOKEN_ERROR)

ROLE_BASE_INFO = _reflection.GeneratedProtocolMessageType('ROLE_BASE_INFO', (_message.Message,), dict(
  DESCRIPTOR = _ROLE_BASE_INFO,
  __module__ = 'gateway_pb2'
  # @@protoc_insertion_point(class_scope:protocol.ROLE_BASE_INFO)
  ))
_sym_db.RegisterMessage(ROLE_BASE_INFO)

MSG_GATEWAY_VERIFY_TOKEN = _reflection.GeneratedProtocolMessageType('MSG_GATEWAY_VERIFY_TOKEN', (_message.Message,), dict(
  DESCRIPTOR = _MSG_GATEWAY_VERIFY_TOKEN,
  __module__ = 'gateway_pb2'
  # @@protoc_insertion_point(class_scope:protocol.MSG_GATEWAY_VERIFY_TOKEN)
  ))
_sym_db.RegisterMessage(MSG_GATEWAY_VERIFY_TOKEN)

MSG_GATEWAY_VERIFY_TOKEN_RESULT = _reflection.GeneratedProtocolMessageType('MSG_GATEWAY_VERIFY_TOKEN_RESULT', (_message.Message,), dict(
  DESCRIPTOR = _MSG_GATEWAY_VERIFY_TOKEN_RESULT,
  __module__ = 'gateway_pb2'
  # @@protoc_insertion_point(class_scope:protocol.MSG_GATEWAY_VERIFY_TOKEN_RESULT)
  ))
_sym_db.RegisterMessage(MSG_GATEWAY_VERIFY_TOKEN_RESULT)


# @@protoc_insertion_point(module_scope)
