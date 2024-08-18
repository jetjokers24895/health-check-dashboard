
const createService = ({name, url}) => {
  fetch("/api/services", {
    method: "POST",
    body: JSON.stringify({
      name: name,
      url: url,
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
    url: $("#newServiceURL").val(),
  });
});