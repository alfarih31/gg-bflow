package message_adapters

import (
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/nb-go-keyvalue"
)

func GRPCKeyValueToKeyValue(kv []*gg_bflow.KeyValue) keyvalue.KeyValue {
	oKv := keyvalue.KeyValue{}

	for _, v := range kv {
		switch v.Type {
		case gg_bflow.ValueType_STR:
			oKv[v.Key] = v.StrValue
		case gg_bflow.ValueType_INT:
			oKv[v.Key] = v.IntValue
		case gg_bflow.ValueType_FLOAT:
			oKv[v.Key] = v.FloatValue
		case gg_bflow.ValueType_BOOL:
			oKv[v.Key] = v.BoolValue
		}
	}

	return oKv
}

func KeyValueToGRPCKeyValue(kv keyvalue.KeyValue) []*gg_bflow.KeyValue {
	gKeyVal := []*gg_bflow.KeyValue{}
	for k, v := range kv {
		d := &gg_bflow.KeyValue{
			Key: k,
		}

		switch _v := v.(type) {
		case string:
			d.StrValue = &_v
			d.Type = gg_bflow.ValueType_STR
		case int64:
			d.IntValue = &_v
			d.Type = gg_bflow.ValueType_INT
		case int:
			_v64 := int64(_v)
			d.IntValue = &_v64
			d.Type = gg_bflow.ValueType_INT
		case bool:
			d.BoolValue = &_v
			d.Type = gg_bflow.ValueType_BOOL
		case float32:
			d.FloatValue = &_v
			d.Type = gg_bflow.ValueType_FLOAT
		case float64:
			_v32 := float32(_v)
			d.FloatValue = &_v32
			d.Type = gg_bflow.ValueType_FLOAT
		}

		gKeyVal = append(gKeyVal, d)
	}

	return gKeyVal
}
