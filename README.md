```
                   ___      ___   ___   ______    __        ______   .___________. __
                  /   \     \  \ /  /  /  __  \  |  |      /  __  \  |           ||  |
                 /  ^  \     \  V  /  |  |  |  | |  |     |  |  |  | `---|  |----`|  |
                /  /_\  \     >   <   |  |  |  | |  |     |  |  |  |     |  |     |  |
               /  _____  \   /  .  \  |  `--'  | |  `----.|  `--'  |     |  |     |  `----.
              /__/     \__\ /__/ \__\  \______/  |_______| \______/      |__|     |_______|
```

![](https://img.shields.io/twitter/follow/BukanYahya?style=social)
![](https://img.shields.io/github/go-mod/go-version/k1m0ch1/axolotl)
![](https://img.shields.io/github/v/release/k1m0ch1/axolotl)
![](https://img.shields.io/github/commit-activity/w/k1m0ch1/axolotl)
![](https://img.shields.io/github/last-commit/k1m0ch1/axolotl)
![](https://img.shields.io/github/release-date/k1m0ch1/axolotl)

# Axolotl - ez vuln record

axolotl is a pentest collaboration tools, comes with a simple feature, and it want to keep it simple, you only need to install axolotl and git on your machine. It has a main purpose to store and collaborate all finding with your team or yourself, and axolotl process the data to simplify lookup data, make a simple statistic and generate a simple report.

<p align="center">
    <img height="50%" width="50%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/preview-1.png">
</p>

When it comes with pentestration collaboration tools, It becomes hard when you manage the document based, sometime rely on file you store on harddrive or cloud storage is hard to manage, and you need times to makes a report or statistic.


Another option, you can manage every finding with "any" pentest documentation tools, sometime with great feature generate documentation and statistic, but it comes with problem you need to pay, sometime you need to install on your server/local and have many requirement to install.


axolotl comes with a simple feature, and it want to keep it simple, you only need to install axolotl and git on your machine. It has a main purpose to store and collaborate all finding with your team or yourself, and axolotl process the data to simplify lookup data, make a simple statistic and generate a simple report.

Axolotl inspired from [nuclei](https://github.com/projectdiscovery/nuclei) project, where I'm using nuclei as the collaboration tools for poc.

# tl;dr axolotl

```
!!Attention!! All data at the screenshot is all dummy, not real data
```

1. Download the binary from [Release](https://github.com/k1m0ch1/axolotl/releases)
2. Install on your machine
3. Run `axolotl init` to create new directory structure
4. Generate host identity and input as you needs (if you didn't need the key, just delete the key)

```
axolotl add -d domain.com
```
<p align="center">
    <img height="60%" width="60%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/img/created-domain.png">
</p>
5. Generate Vulnerability Finding and input as you needs (if you didn't need the key, just delete the key)

```
axolotl add -d domain.com -v vuln-name-without-space
```
<p align="center">
    <img height="60%" width="60%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/img/vuln-add.png">
</p>
6. List all current Host

```
axolotl lookup host
```
<p align="center">
    <img height="50%" width="50%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/img/lookup-host.png">
</p>
7. List all current Vuln

```
axolotl lookup vuln
```
<p align="center">
    <img height="50%" width="50%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/img/lookup-vuln.png">
</p>
8. Information Host with Vuln

```
axolotl info -d domain.com
```
<p align="center">
    <img height="65%" width="65%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/img/info-domain.png">
</p>
9. simple statistic about your finding

```
axolotl stat
```

<p align="center">
    <img height="65%" width="65%" src="https://github.com/k1m0ch1/axolotl/raw/master/.github/img/statistic.png">
</p>

10. repeat from `4` to add more host and vuln finding

Check [How to use](https://axolotl.readthedocs.io/en/latest/) page for detail how to use


## Release and Contributing

We appreciate all contributions. If you are planning to contribute any bug-fixes, please do so without further discussions.

If you plan to contribute new features, new tuners, new training services, etc. please first open an issue or reuse an exisiting issue, and discuss the feature with us. We will discuss with you on the issue timely or set up conference calls if needed.

To learn more about making a contribution to axolotl, please refer to our How-to contribution page.

Please let us know if you encounter a bug by filling an issue.

We appreciate all contributions and thank all the contributors!

<a href="https://github.com/k1m0ch1/axolotl/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=k1m0ch1/axolotl" />
</a>


