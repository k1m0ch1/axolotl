```
     ___      ___   ___   ______    __        ______   .___________. __
    /   \     \  \ /  /  /  __  \  |  |      /  __  \  |           ||  |
   /  ^  \     \  V  /  |  |  |  | |  |     |  |  |  | `---|  |----`|  |
  /  /_\  \     >   <   |  |  |  | |  |     |  |  |  |     |  |     |  |
 /  _____  \   /  .  \  |  `--'  | |  `----.|  `--'  |     |  |     |  `----.
/__/     \__\ /__/ \__\  \______/  |_______| \______/      |__|     |_______|
```

A simple bug reporting tools for hackers to input the finding and Host Identity, by record all the finding or host with git and without needs to install the tools.

File Structure  
```
root of this repo
├── hosts
| ├── all.txt
| ├── Platform&Data
| | ├── platformdata.txt
| | ├── Infrastructure
| | |  ├── infrastructure.txt
| | |  └── apps
| | |   ├── vulnerabilty-name.vuln
| | |   └── poc
├── report-gen
| └── src
├── host-gen
| └── src
├── axolotl
| └── src
└── Makefile
```

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
Vulnerability Report Format
using YAML nuclei format
```
id: url-slug-name-vuln-with-version

info:
  finding-name: Name of the Vulnerability
  url: 
  author: person1, person2
  email: person1@efishery.com, person2@efishery.com
  severity: high
  cvss-score-vector: CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:N
  owasp-score-vector: (SL:0/M:0/O:0/S:0/ED:0/EE:0/A:0/ID:0/LC:0/LI:0/LAV:0/LAC:0/FD:0/RD:0/NC:0/PV:0)
  tags: tag,vulnerability,list
  report: default-with-file-name.md 

poc:
  - path: https://target.com/??adad=asdasd
    step-to-reproduce:
      - desc: first step
        image: file.png
      - desc: second step
        image: file.png
    image: file.png
    nuclei: file.template
    exploit: exploit.sh
  - path: https://target.com/??adad=asdasd
    step-to-reproduce:
      - desc: first step
        image: file.png
      - desc: second step
        image: file.png
    image: file.png
    nuclei: file.template
    exploit: exploit.sh

recommendation-to-fix:
  - info:
    desc:
    link:
    step-to-fix:
      - desc: 
        image:
    image:

status: 
  requested_at:
  open_at:
  reviewed_at:
  approved_at:
  fixed_at:
  validated_at:
  duplidated_at:
  hold_at:
  rejected_at:
  closed_at:
  complete_at:
```