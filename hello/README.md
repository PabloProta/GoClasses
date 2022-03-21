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


# Compile and Install App
  From the hello directory run:
  <pre>go build</pre> and execute <code>./hell.exo</code> to check if executable works.
  You may need to install the app to execute from any directory, so to do this you need to discover the install path with:
  <pre>go list -f '{{.Target}}'</pre>
  So then you need to Add the Go install directory to your system's shell path:
  <pre>set PATH=%PATH%;C:\Users\Pablo AP\Desktop\Google\Go\bin\hello.exe</pre>
  
  As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:
  <pre> go env -w GOBIN=/path/to/your/bin </pre> or <pre>go env -w GOBIN=C:\path\to\your\bin </pre>

  :warning:Actually any command work for me, i had to add the ambient variable manually;

  
