
const createService = ({name, command}) => {
  fetch("/api/services", {
    method: "POST",
    body: JSON.stringify({
      name: name,
      command: command,
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then(response => {
    if (response.ok) {
      window.location.href = "/";
    }
  }).catch(err => console.log(err));
}
$("#newServiceBtn").on("click", function (e) {
  e.preventDefault();
  createService({
    name: $("#newServiceName").val(),
    command: $("#newServiceCommand").val(),
  });
});