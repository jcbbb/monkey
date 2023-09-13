let output = document.querySelector(".output")
let form = document.getElementById("form")
let select = document.querySelector(".header__select");
let editable = document.querySelector(".editable")

let examples = {
  "hello-world": `"Hello world!";`,
  "functions": `let add = fn(a, b) {
  return a + b;
};

let multiply = fn(a, b) {
  return a * b;
};

let hashmap = { "two": 2 };
let array = [10, 90, 200];

let double = fn(x) { multiply(x, hashmap.two) };

let map = fn(arr, f) {
  let iter = fn(arr, acc) {
    if (len(arr) == 0) {
      acc
    } else {
      iter(rest(arr), push(acc, f(first(arr))))
    }
  };

  iter(arr, [])
};

map(array, double);
`
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
