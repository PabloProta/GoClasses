**TIPS** :wave:
-

You may need to edit module from hello to recognize the greetings module to do this its just needed do:
<br><code>go mod edit -replace \<name of module>=\<folder of module></code>

In This Tutorial i used:
<br><code>go mod edit -replace example/greetings=../greetings</code>

### :warning: Dont Forget to enter in "hello" folder before to do this.  
  And after doing this you need to uptade thee hello module with <code>go mod tidy</code>.<br>
  You can see after run this command in go.mod that included a require directive with a <i>pseudo-version number</i> since the greetings is local module instead a remote module that a have a version number.
      
For more about version number [see](https://go.dev/doc/modules/version-numbers)