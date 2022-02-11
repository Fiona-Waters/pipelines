# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import kfp_server_api
from kfp_server_api.models.protobuf_struct import ProtobufStruct  # noqa: E501
from kfp_server_api.rest import ApiException

class TestProtobufStruct(unittest.TestCase):
    """ProtobufStruct unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ProtobufStruct
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kfp_server_api.models.protobuf_struct.ProtobufStruct()  # noqa: E501
        if include_optional :
            return ProtobufStruct(
                fields = {
                    'key' : kfp_server_api.models.protobuf_value.protobufValue(
                        null_value = 'NULL_VALUE', 
                        number_value = 1.337, 
                        string_value = '0', 
                        bool_value = True, 
                        struct_value = kfp_server_api.models.protobuf_struct.protobufStruct(), 
                        list_value = kfp_server_api.models.protobuf_list_value.protobufListValue(
                            values = [
                                kfp_server_api.models.protobuf_value.protobufValue(
                                    number_value = 1.337, 
                                    string_value = '0', 
                                    bool_value = True, )
                                ], ), )
                    }
            )
        else :
            return ProtobufStruct(
        )

    def testProtobufStruct(self):
        """Test ProtobufStruct"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()