
function createPrompt() {
  const html = `
    <div
      class="flex justify-start items-center"
    >
      <div class="font-bold text-yellow-400 mr-3 text-nowrap">juan@portfolio:~ $</div>
      <div class="w-fit flex justify-start items-center">
        <div id="command"></div>
        <div id="cursor" class="bg-[#f8f8f2] w-2 h-4 animate-blink"></div>
      </div>
    </div>

    <div id="info" hx-on::after-swap="this.removeAttribute('id'); appendPrompt();"></div>
  `;

  const doc = new DOMParser().parseFromString(html, "text/html");
  const prompt = document.createDocumentFragment();

  while (doc.body.childElementCount > 0) {
    const child = doc.body.firstElementChild;
    prompt.appendChild(child);
    htmx.process(child);
  }

  return prompt;
}

function appendPrompt() {
  const content = document.getElementById("content");
  const prompt = createPrompt();
  const promptText = prompt.firstElementChild;
  content.appendChild(prompt);
  promptText.scrollIntoView({ behavior: "smooth" });
}

function clearContent() {
  const content = document.getElementById("content");
  content.innerHTML = "";
  appendPrompt();
}
