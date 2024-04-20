## Blueprint
```
type Range struct {
  Min string
  Max string
}

func (r Range) IntR() (min, max int)
func (r Range) FloatR() (min, max float64)
func (r Range) DateR() (min, max time.Time)
```
* Range is for option of user to get data between the min and max.
* Since Range holds the min & max as string, IntR converts to int for int type, FloatR to float and DateR to time.Time
```
type Prop struct {
  Id string
  Name string
  Type string
  Length int
  In []string
  Range Range
  Unique bool
  Prefix string
  Suffix string
}
func (p Prop) GetInInt() (arr []int)
func (p Prop) GetInFloat() (arr []float64)
func (p Prop) GetInDate() (arr []time.Time)
func (p Prop) GetPrefixInt() (prefix int)
func (p Prop) GetPrefixFloat() (prefix float64)
func (p Prop) GetSuffixInt() (prefix int)
func (p Prop) GetSuffixFloat() (prefix float64)
```
* Prop defines the actual schematics of a single prop
* Since the 'In' property holds array of string, GetInInt, GetInFloat and GetInDate converts the array to the specific type
* Same goes for the prefix and suffix methods
```
type Blueprint struct {
  Volume int
  Props []Prop
}
func (b Blueprint) IsValidRangeInt() bool
func (b Blueprint) IsValidRangeDate() bool
```
* Blueprint defines the actual schematics of the entire json object
* IsValidRangeInt and IsValidRangeDate checks if the total number of data that can be derived from the range is higher the Blueprint.Volume for props that have the Unique property as true
## ProcessedProp
```
type ProcessedProp struct {
  Id string
  Name string
  Value interface{}
}
func (pp ProcessedProp) Addr() (position int, classes []int)
```
* This is to hold the data after it has been processed
