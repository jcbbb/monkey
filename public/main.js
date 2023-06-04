let output = document.querySelector(".output")
let form = document.getElementById("form")
let select = document.querySelector(".header__select");
let editable = document.querySelector(".editable")

let examples = {
  "hello-world": `"Hello world!";`,
  "functions": `let add = fn(a, b) {\n  return a + b;\n};\n\nadd(34, 35);\n`,
}

select?.addEventListener("change", (e) => {
  let example = examples[e.target.value]
  editable.textContent = example;
})

form?.addEventListener("submit", async (e) => {
  e.preventDefault()
  let data = new FormData(e.target)

  let response = await fetch(e.target.action, {
    method: "POST",
    body: data.get("body")
  })

  output.classList.toggle("error", !response.ok)

  let result = await response.text()
  output.textContent = result
})
