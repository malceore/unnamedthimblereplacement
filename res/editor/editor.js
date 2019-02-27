const textarea = document.querySelector(".editor");
const iframe = document.querySelector("iframe");
const btn = document.querySelector("button");

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


var editor = CodeMirror.fromTextArea(textarea, {
  theme: 'cobalt',
  lineNumbers: true,
  mode : "xml",
  htmlMode: true
});

editor.setSize(null, "100%");