Host Identity Format
```
target:
  url: domain or apps 
  tech-stack: js, lalala
  open-port: 80/http, 81/ssh
  filtered-port: 
  host-ip: 
  recon:
    - tools: nmap
      report: file.html
    - tools: theHarvester
      report: file.txt
  vuln-scan:
    - tools: burp
      report: lala.html

recon-found:
  source-code-leak:
    - info: name of found
      desc: lalala
      report: file.html
      image: file.png
  GHDB:
    - info: n/a
      desc: lalala
      report: file.html
      image: file.png
  virus-total:
    - info:
      detection-result: file
      details-result: file
      desc: 
  the-harvester:
    - info:
      desc:
      result: file
  web-archive:
    - info:
      desc:
      image: 
  defacement:
  osint-discovery:
  directory-index:
  reverse-ip-check:
  project-management-tools:
    - info:
      desc:

behaviour:
  unique-behaviour:
    - info:
      desc:
    - info:
      desc:
  third-party-hosted-content:
    - info:
      desc:

http-responses:
  - result:
    desc:

ssl-implemented:
  - result:
    http-redirect: (True/ False)
    desc:

shodan-result:
  - info:
    result-ip:https://www.shodan.io/search?query=net:1.2.3.4,5.6.7.8,9.10.11.12
```