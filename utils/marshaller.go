//Copyright 2016-2017 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package utils

import (
	"encoding/json"

	"errors"
)

/*
MarshalB returns the JSON encoding of message as byte array. If the message cannot be encoded, an error is returned with an empty byte array. Hence no nil value is returned.
*/
func MarshalB(message interface{}) ([]byte, error) {
	b, err := json.Marshal(message)
	if err == nil {
		return b, nil
	}
	return []byte{}, err

}

/*
MarshalS returns the JSON encoding of message as string. If the message cannot be encoded, an error is returned with an empty byte array. Hence no nil value is returned.
*/
func MarshalS(message interface{}) (string, error) {
	b, err := MarshalB(message)
	return string(b), err
}

/*
MarshalSD is comparable to MarshalS. However, it returns always a valid string and no error. When an error occurs, an empty string is returned.
*/
func MarshalSD(message interface{}) string {
	result, err := MarshalS(message)
	if err != nil {
		result = "{}"
	}
	return result
}

/*
Unmarshal stores the values of a Json encoded message 'orig' in a data type 'result'.
*/
func Unmarshal(orig interface{}, result interface{}) (err error) {

	switch v := orig.(type) {
	case string:
		err = unmarshalB([]byte(v), result)
	case []byte:
		err = unmarshalB(v, result)
	default:
		err = errors.New("Could not unmarshal value")
	}

	return err
}

func unmarshalB(message []byte, result interface{}) error {
	if err := json.Unmarshal(message, result); err != nil {
		return err
	}
	return nil
}
