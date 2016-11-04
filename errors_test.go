package gencheck

//
// import (
// 	"encoding/json"
// 	"errors"
// 	"testing"
//
// 	"github.com/stretchr/testify/suite"
// )
//
// // NoValidate
// type NoValidate struct{}
//
// // HasValidate
// type HasValidate struct{}
//
// // Validate
// func (s HasValidate) Validate() error {
// 	return errors.New("Validating 'HasValidate' instance")
// }
//
// // gencheckErrorsTestSuite
// type gencheckErrorsTestSuite struct {
// 	suite.Suite
// }
//
// // TestgencheckTypesTestSuite
// func TestgencheckTypesTestSuite(t *testing.T) {
// 	suite.Run(t, new(gencheckErrorsTestSuite))
// }
//
// // TestValidate
// func (s *gencheckErrorsTestSuite) TestValidate() {
// 	a := HasValidate{}
// 	b := &HasValidate{}
// 	c := NoValidate{}
//
// 	err := Validate(a)
// 	s.Equal(errors.New("Validating 'HasValidate' instance"), err)
//
// 	err = Validate(b)
// 	s.Equal(errors.New("Validating 'HasValidate' instance"), err)
//
// 	err = Validate(c)
// 	s.Nil(err)
// }
//
// // TestErrorSliceError_Empty
// func (s *gencheckErrorsTestSuite) TestErrorSliceError_Empty() {
// 	ea := gencheck.ErrorSlice{}
//
// 	s.Equal("[]", ea.Error())
// }
//
// // TestErrorSliceError_MultiElements
// func (s *gencheckErrorsTestSuite) TestErrorSliceError_MultiElements() {
// 	ea := gencheck.ErrorSlice{
// 		errors.New("foo"),
// 		errors.New("bar"),
// 		nil,
// 		gencheck.ErrorSlice{errors.New("this is"), errors.New("nested")},
// 	}
//
// 	s.Equal("[\"foo\",\"bar\",null,[\"this is\",\"nested\"]]", ea.Error())
// }
//
// // TestErrorMapError_Empty
// func (s *gencheckErrorsTestSuite) TestErrorMapError_Empty() {
// 	em := gencheck.ErrorMap{}
// 	s.Equal("{}", em.Error())
// }
//
// // TestErrorMapError_NilValue
// func (s *gencheckErrorsTestSuite) TestErrorMapError_NilValue() {
// 	em := gencheck.ErrorMap{
// 		"flat":                nil,
// 		"nestedErrorSlice":    gencheck.ErrorSlice{errors.New("this is"), errors.New("nested")},
// 		"nestedEmptyErrorMap": make(gencheck.ErrorMap),
// 	}
//
// 	expectedJSONAsMap := make(map[string]interface{})
// 	actualJSONAsMap := make(map[string]interface{})
// 	json.Unmarshal([]byte(`{"flat": null,"nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
// 	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)
//
// 	s.Equal(expectedJSONAsMap, actualJSONAsMap)
// }
//
// // TestErrorMapError_MultipleValues
// func (s *gencheckErrorsTestSuite) TestErrorMapError_MultipleValues() {
// 	em := gencheck.ErrorMap{
// 		"flat":                errors.New(`"flat" "error"`),
// 		"nestedErrorSlice":    gencheck.ErrorSlice{errors.New("this is"), errors.New("nested")},
// 		"nestedEmptyErrorMap": make(gencheck.ErrorMap),
// 	}
//
// 	expectedJSONAsMap := make(map[string]interface{})
// 	actualJSONAsMap := make(map[string]interface{})
// 	json.Unmarshal([]byte(`{"flat": "\"flat\" \"error\"","nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
// 	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)
//
// 	s.Equal(expectedJSONAsMap, actualJSONAsMap)
// }
