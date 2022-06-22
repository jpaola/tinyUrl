const axios = require('axios').default;

axios.post("https://cleanuri.com/api/v1/shorten", {
    url: "https://marketplace.digitalocean.com/"
  })
  .then(function (response) {
    console.log(response.data);
  })
  .catch(function (error) {
    console.log("There was an error with this request: ", error);
  });