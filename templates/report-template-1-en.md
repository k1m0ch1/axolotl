---
title: "Pentest Report for {{.Host.Info.Desc}}"
author: ["{{.User.ProjectOwner}}", "{{.User.Email}}"]
date: "2022-03-25"
subject: "Markdown"
keywords: [Markdown, Example]
subtitle: "{{.Host.Info.URL}}"
lang: "en"
titlepage: true
titlepage-color: "1E90FF"
titlepage-text-color: "FFFAFA"
titlepage-rule-color: "FFFAFA"
titlepage-rule-height: 2
book: true
classoption: oneside
code-block-font-size: \scriptsize
---
# Pentest Report for {{.Host.Info.Desc}}

## Introduction

Examle of the reports contain all finding from {{.Host.Info.Desc}} with IP {{.Host.Info.HostIP}} in order to inform the existing vulnerability, as in written on Bug Bounty Program gojek.com.

My Name is {{.User.ProjectOwner}}, I'm the Cyber Security Researcher, you can contact me from email {{.User.Email}}

## Requirements

In order to understand about this report, you need to understand a very basic CVSS or cyber security finding.

# High-Level Summary

Our assessment found some of the risk on your system/app and here is the summary:

We found --VULNVOUNT-- Vulnerability on you domain, and here is the detail:


## Recommendations

I recommend to patch the vulnerability from the highest to the lowest vulnerability, since the highest risk vulnerability is much more easy to gain access it will be the best to understand about the vulnerability and follow the recommendation at the detail recommendation

# Methodologies

This document is used the common and standarized methodologies to follow the best practices, with this document we much more follow the best practice from OWASP Top 10, This best practice is much more updated every 2-3 years, and many people likely to use this as the best practice for vulnerability assessment.

## Information Gathering

The information gathering portion of a penetration test focuses on identifying the scope of the penetration test. The specific IP addresses were:

- 192.168.

# Vulnerability Finding

{{ range .Vulns }}
## {{ .VulnInfo.FindingName }}

**Description**

{{ .VulnInfo.Desc }}


Severity  |CVSS-Vector|
----------|---|
{{ .VulnInfo.RiskRatings.Severity }}      | {{ .VulnInfo.RiskRatings.CVSS }}  |

**Vulnerability Type: {{ .VulnInfo.VulnType }}**

**Affected URL: {{ .VulnInfo.URL }} **

### Steps to reproduce

{{ range .ProofOfConcept }}

**Description**

{{ .Desc }}

{{ range .Images }}
![{{.Caption}}]({{.Path}})
{{ end }}

**Affected URL: {{ .Path }} **
{{ range $index, $val := .StepsToReproduce }}
{{ add $index 1 }}. {{ $val.Desc }}
{{ range $val.Images }}
![{{.Caption}}]({{.Path}})
{{ end }}
{{ end }}

**Payload**

```
{{ .Payload }}
```

{{ end }}

### Recommendation

{{ end }}