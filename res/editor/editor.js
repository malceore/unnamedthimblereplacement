const textarea = document.querySelector(".editor");
const iframe = document.querySelector("iframe");
/*const btn = document.querySelector("button");

btn.addEventListener("click", () => {
  var html = editor.getValue();
  iframe.src = "data:text/html;charset=utf-8," + encodeURI(html);
});

textarea.addEventListener('keyup',()=>{
  var html = editor.getValue();
  iframe.src = "data:text/html;charset=utf-8," + encodeURI(html);
})

textarea.addEventListener("paste", function(e) {
  e.preventDefault();
  var text = e.clipboardData.getData("text/plain");
  document.execCommand("insertText", false, text);
});

*/

/*
** Loads in files needed for the website, will be placed in at runtime by the Golang service when it wants to load the code for a website.
**
*/
function loadFiles(filearr){
  let filediv = document.querySelector(".filediv");
  for (var i=0; i<filearr.length; i+=2){
    var inputElement = document.createElement('input');
    inputElement.type = "button";
    inputElement.value = filearr[i];
    inputElement.contents = filearr[i+1];
    inputElement.addEventListener('click', function(){
      loadContents("" + this.contents);
    });
    filediv.appendChild(inputElement);
  }
  console.log("Loading files Complete!");
}

function updateContents(){
  let cont = editor.getValue();
  iframe.src = "data:text/html;charset=utf-8," + encodeURI(cont);
}

function saveContents(){
  //let cont = editor.getValue(contents);
  //iframe.src = "data:text/html;charset=utf-8," + encodeURI(cont);
  console.log("Saving contents!");
}

function loadContents(contents){
  console.log(contents);
  editor.setValue(contents);
  iframe.src = "data:text/html;charset=utf-8," + encodeURI(contents);
}

var editor = CodeMirror.fromTextArea(textarea, {
  theme: 'cobalt',
  lineNumbers: true,
  mode : "xml",
  htmlMode: true
});

editor.setSize(null, "100%");
