# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: networking/v1alpha3/workload_group.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='networking/v1alpha3/workload_group.proto',
  package='istio.networking.v1alpha3',
  syntax='proto3',
  serialized_options=_b('Z istio.io/api/networking/v1alpha3'),
  serialized_pb=_b('\n(networking/v1alpha3/workload_group.proto\x12\x19istio.networking.v1alpha3\x1a\x1fgoogle/api/field_behavior.proto\"\xa0\x02\n\rWorkloadGroup\x12\x42\n\x05ports\x18\x01 \x03(\x0b\x32\x33.istio.networking.v1alpha3.WorkloadGroup.PortsEntry\x12\x44\n\x06labels\x18\x02 \x03(\x0b\x32\x34.istio.networking.v1alpha3.WorkloadGroup.LabelsEntry\x12\x0f\n\x07network\x18\x03 \x01(\t\x12\x17\n\x0fservice_account\x18\x04 \x01(\t\x1a,\n\nPortsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\"Z istio.io/api/networking/v1alpha3b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,])




_WORKLOADGROUP_PORTSENTRY = _descriptor.Descriptor(
  name='PortsEntry',
  full_name='istio.networking.v1alpha3.WorkloadGroup.PortsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1alpha3.WorkloadGroup.PortsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1alpha3.WorkloadGroup.PortsEntry.value', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=302,
  serialized_end=346,
)

_WORKLOADGROUP_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='istio.networking.v1alpha3.WorkloadGroup.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1alpha3.WorkloadGroup.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1alpha3.WorkloadGroup.LabelsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=348,
  serialized_end=393,
)

_WORKLOADGROUP = _descriptor.Descriptor(
  name='WorkloadGroup',
  full_name='istio.networking.v1alpha3.WorkloadGroup',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ports', full_name='istio.networking.v1alpha3.WorkloadGroup.ports', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='labels', full_name='istio.networking.v1alpha3.WorkloadGroup.labels', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='network', full_name='istio.networking.v1alpha3.WorkloadGroup.network', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='service_account', full_name='istio.networking.v1alpha3.WorkloadGroup.service_account', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_WORKLOADGROUP_PORTSENTRY, _WORKLOADGROUP_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=105,
  serialized_end=393,
)

_WORKLOADGROUP_PORTSENTRY.containing_type = _WORKLOADGROUP
_WORKLOADGROUP_LABELSENTRY.containing_type = _WORKLOADGROUP
_WORKLOADGROUP.fields_by_name['ports'].message_type = _WORKLOADGROUP_PORTSENTRY
_WORKLOADGROUP.fields_by_name['labels'].message_type = _WORKLOADGROUP_LABELSENTRY
DESCRIPTOR.message_types_by_name['WorkloadGroup'] = _WORKLOADGROUP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WorkloadGroup = _reflection.GeneratedProtocolMessageType('WorkloadGroup', (_message.Message,), {

  'PortsEntry' : _reflection.GeneratedProtocolMessageType('PortsEntry', (_message.Message,), {
    'DESCRIPTOR' : _WORKLOADGROUP_PORTSENTRY,
    '__module__' : 'networking.v1alpha3.workload_group_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.WorkloadGroup.PortsEntry)
    })
  ,

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _WORKLOADGROUP_LABELSENTRY,
    '__module__' : 'networking.v1alpha3.workload_group_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.WorkloadGroup.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _WORKLOADGROUP,
  '__module__' : 'networking.v1alpha3.workload_group_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.WorkloadGroup)
  })
_sym_db.RegisterMessage(WorkloadGroup)
_sym_db.RegisterMessage(WorkloadGroup.PortsEntry)
_sym_db.RegisterMessage(WorkloadGroup.LabelsEntry)


DESCRIPTOR._options = None
_WORKLOADGROUP_PORTSENTRY._options = None
_WORKLOADGROUP_LABELSENTRY._options = None
# @@protoc_insertion_point(module_scope)
