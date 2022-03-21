**TIPS** :wave:
-

## Function Structure:

![function](https://go.dev/doc/tutorial/images/function-syntax.png)

## Declaring Variable
<pre>
    Name      Value 
      v         v
    message:= "Hello World!" </pre>
  Or
  <pre>
    var message string
    message= "Hello World!" </pre>

## Conditions Structures 
<pre>
  if name == "Prota"{
    //if return true
  }</pre>

As you can see in go "if" statement dont need the parenteses around the comparison.

## Map Structure
To initialize a map :
<pre>
 variableName := make(map[key-type]value-type)
</pre>

## For Structure
<pre>
	for i := 0; i < count; i++ {
		//code...
	}</pre>
  or     
  <pre>
  for index, name:= range names {
          //code...
	}</pre>

  In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value

# Null values
  For null values you put <code>nil</code> instead null

# Test
  If you want improve tests you may need to create a teste file.  In my case i put <code>greetings_test.go</code> the <code>_test.go</code> tells to the <code>go test</code>command that this file contains test functions.
  <br>Test function names have the form <code><b>Test</b>Name</code>, where <i>Name</i> says something about the specific test.  Also test functions take a pointer to the testing package's testing.T [type](https://pkg.go.dev/testing/#T) as a parameter.  You use this parameter's methods for reporting and logging from your test.
