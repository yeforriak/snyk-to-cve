# snyk-to-cve
Consumes the output from `snyk-cli` or `docker scan` attaching CVEs

## Usage
Building and running as command:
```
go build .
./snyk-to-cve snyk-output.txt
```

## Example
```
docker scan ubuntu-impish > snyk-output.txt
./snyk-to-cve snyk-output.txt
https://snyk.io/vuln/SNYK-UBUNTU2110-TAR-1744334 -> CVE-2019-9923
https://snyk.io/vuln/SNYK-UBUNTU2110-SHADOW-1758374 -> CVE-2013-4235
https://snyk.io/vuln/SNYK-UBUNTU2110-PCRE3-1747092 -> CVE-2019-20838
https://snyk.io/vuln/SNYK-UBUNTU2110-PCRE3-1755307 -> CVE-2017-11164
https://snyk.io/vuln/SNYK-UBUNTU2110-LIBTASN16-1752539 -> CVE-2018-1000654
https://snyk.io/vuln/SNYK-UBUNTU2110-LIBSEPOL-1735893 -> CVE-2021-36087
https://snyk.io/vuln/SNYK-UBUNTU2110-LIBSEPOL-1735898 -> CVE-2021-36085
https://snyk.io/vuln/SNYK-UBUNTU2110-LIBSEPOL-1735900 -> CVE-2021-36086
https://snyk.io/vuln/SNYK-UBUNTU2110-LIBSEPOL-1735902 -> CVE-2021-36084
https://snyk.io/vuln/SNYK-UBUNTU2110-KRB5-1749300 -> CVE-2018-5709
https://snyk.io/vuln/SNYK-UBUNTU2110-GMP-1921286 -> CVE-2021-43618
https://snyk.io/vuln/SNYK-UBUNTU2110-COREUTILS-1756916 -> CVE-2016-2781
https://snyk.io/vuln/SNYK-UBUNTU2110-ZLIB-2433596 -> CVE-2018-25032

```
