const textarea = document.querySelector(".editor");
const iframe = document.querySelector("iframe");
var currentFile;

/*
** Loads in files needed for the website, will be placed in at runtime by the Golang service when it wants to load the code for a website.
**
*/
function loadFiles(filearr){
  let filediv = document.querySelector(".filediv");
  var first = true;
  for (var i=0; i<filearr.length; i+=2){

    var inputElement = document.createElement('input');
    inputElement.type = "button";
    inputElement.value = filearr[i];
    inputElement.contents = filearr[i+1];
    inputElement.addEventListener('click', function(){
      loadContents("" + atob(this.contents));
      currentFile = this;
    });
    filediv.appendChild(inputElement);

    if(first){
      first = false;
      loadContents(atob(filearr[i+1]));
      currentFile = inputElement;
    }

  }
  console.log("Loading files Complete!");
}

function updateContents(){
  let cont = editor.getValue();
  iframe.src = "data:text/html;charset=utf-8," + encodeURI(cont);
  saveContents();
}

function saveContents(){
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", "/save", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  // Need to encode as Base64 as it helps keep spacing without worry, then re-encode as
  //	& need to be encoded for URI or we get weird behavior.
  xhttp.send("&fileid=" + currentFile.value + "&contents=" + btoa(currentFile.contents).replace('&', '%26'));
  currentFile.contents = editor.getValue();
  //console.log("DEBUG::Saving contents! ");//\n Debug: current contents:" + currentFile.contents + " fileid::" + currentFile.value);
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

/*
editor.addEventListener('keyup',()=>{
  updateContents();
})*/
