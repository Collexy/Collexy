package globals

import (
  //"errors"
  //"collexy/helpers"
  "strings"
  "strconv"
  "fmt"
  "database/sql/driver"
)

type IntSlice []int 

func (b *IntSlice) Scan(src interface{}) error { 
        switch src := src.(type) { 
        case nil: 
                *b = nil 
                return nil 

        case []byte: 
                // TODO: parse src into *b
          var intArr []int
          intArrString := string(src)
          intArrString = strings.Replace(intArrString, "{", "", -1)
          intArrString = strings.Replace(intArrString, "}", "", -1)
          var lol []string
          lol = strings.Split(intArrString, ",")
          for i := 0; i < len(lol); i++ {
            someval, _ := strconv.Atoi(lol[i])
             intArr = append(intArr, someval)
          }
          *b = intArr

        default: 
                return fmt.Errorf(`unsupported driver -> Scan pair: %T -> *IntSlice`, src) 
        }
        return nil
}

func (b IntSlice) Value() (driver.Value, error) {
  var str string = "{"
  var myarr []int = b
  fmt.Println("driver.Value 1: ")
  fmt.Println(b)
  for i := 0; i < len(myarr); i++ {
    str = str + strconv.Itoa(myarr[i])
    if(i<len(myarr)-1){
      str = str+","
    }
  }
  str = str + "}"
  fmt.Println("driver.Value 2: ")
  fmt.Println(str)
  return str, nil
  //return "{23,24}", nil
  //return "20,21", nil
    // Format b in PostgreSQL's array input format {1,2,3} and return it as as string or []byte.
    // if(b == nil){
    //   return nil, nil
    // } else if(len(*b)>0){
    //   var str string = "{"
    //   for i := 0; i < len(*b); i++ {
    //     str = str + string(*b[i])
    //     if(i<len(b-1)){
    //       str = str+", "
    //     }
    //   }
    //   str = str+"}"
    //   return str
    //   } else {
    //         return fmt.Errorf(`unsupported driver -> Scan pair: %T -> *IntSlice`, src) 
    //   }
    //   return nil
}