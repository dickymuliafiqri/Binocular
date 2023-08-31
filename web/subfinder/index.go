package subfinder

var SubfinderTemplate = `
<!DOCTYPE html>
<html lang="en" class="bg-black px-32 py-5 font-mono text-gray-300">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <script src="https://cdn.tailwindcss.com"></script>
    <title>Binocular | Advanced Subdomain Finder</title>
  </head>
  <body>
    <h1 class="text-[64px] text-indigo-400 font-bold">>{{.PageTitle}}_</h1>
    <table class="font-medium w-full table-auto">
      <tr class="text-left text-semibold">
        <th>Domain</th>
        <th>Port</th>
        <th>Status Code</th>
        <th>Web Server</th>
        <th>Host</th>
        <th>Response Time</th>
      </tr>
      {{ range .Result }}
      <tr
        class='{{if eq .WebServer "cloudflare"}} text-orange-400 {{end}} border-b border-gray-950 border-solid hover:bg-gray-950'
      >
        <td><a href="{{.URL}}" target="_blank">{{.Input}}</a></td>
        <td>{{.Port}}</td>
        <td>{{.StatusCode}}</td>
        <td>{{.WebServer}}</td>
        <td><a href="http://{{.Host}}" target="_blank">{{.Host}}</a></td>
        <td>{{.ResponseTime}}</td>
      </tr>
      {{end}}
    </table>
    <footer class="text-center pt-5">
      <hr class="pt-3 pb-10" />
      <p>Made with ❤ by dickymuliafiqri</p>
    </footer>
  </body>
</html>
`
