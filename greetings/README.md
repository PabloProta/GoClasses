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
