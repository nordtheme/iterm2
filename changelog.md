<p align="center">
  <a href="https://www.nordtheme.com/ports/iterm2" target="_blank">
    <picture>
      <source srcset="https://raw.githubusercontent.com/nordtheme/iterm2/develop/src/assets/nord-iterm2-banner.svg?sanitize=true" width="100%" media="(prefers-color-scheme: light), (prefers-color-scheme: no-preference)" />
      <img srcset="https://raw.githubusercontent.com/nordtheme/iterm2/develop/src/assets/nord-iterm2-banner.svg?sanitize=true" width="100%" />
    </picture>
  </a>
</p>

<p align="center">
  <a href="https://www.nordtheme.com/docs/ports/iterm2" target="_blank">
    <img src="https://img.shields.io/github/release/nordtheme/iterm2.svg?style=flat-square&label=Docs&colorA=4c566a&colorB=88c0d0&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgaGVpZ2h0PSIxNiI%2BCiAgICA8cGF0aCBmaWxsPSIjZDhkZWU5IiBkPSJNMTMuNzQ2IDIuODEzYS42Ny42NyAwIDAgMC0uNTU5LS4xMzNMOCAzLjg0OGwtNS4xODgtMS4xOGEuNjY5LjY2OSAwIDAgMC0uNTcuMTMzLjY3Ny42NzcgMCAwIDAtLjI0Mi41MzF2OC4xMzNjLS4wMDguMzIuMjEuNTk4LjUyLjY2OGw1LjMzMiAxLjE5OWguMjk2bDUuMzMyLTEuMmEuNjY4LjY2OCAwIDAgMCAuNTItLjY2N1YzLjMzMmEuNjU5LjY1OSAwIDAgMC0uMjU0LS41MnpNMy4zMzIgNC4xNjhsNCAuODk4djYuNzY2bC00LS44OTh6bTkuMzM2IDYuNzY2bC00IC44OThWNS4wNjZsNC0uODk4em0wIDAiLz4KPC9zdmc%2BCg%3D%3D"/>
  </a>
</p>

<p align="center">
  <a href="https://github.com/nordtheme/nord/releases/tag/v0.2.0" target="_blank">
    <img src="https://img.shields.io/static/v1.svg?style=flat-square&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABmJLR0QA/wD/AP+gvaeTAAACUklEQVRIibXUQUiTcRjH8e+zvaO0AgPd1mDmNvFVI+ng1UNU1wiDoEIIoUNUdOkQUkER1KWTnRJCLyZBIdEhQvAWQVp5KN9X3bskmHunYbeZbu/TaWXRdOLrc3ue5/f/f/6nP+xwycy8m6y0LGqg0N7UsLB+9jmdbQwZQaNawPBKpCstA3g/7fRCl5k68P7PLNjplRgBQtUAYjmuCriqDP690RqQK8A3WSt1mmZsqbyyM+4pVc5tcvdxYD+W46qVdif/l7Ac95bluGo57piqBqt58e+zaXfSclwNbBQyE+F7KKPAMTuzeGcrQLk2BEREV41dvYAD2jeddrt9BQA6DtYtC3QDKyIM2pmlVl+Bublc2ExGpgS9BuxTLT2byuX2+AYUJTAyPq6GmYwOIPoEOLy7IAO+AYgejTW5dwF2UbiM8kGVs9NO/qo/AGRV5Ybl5E4mEokVT0rdwHdBH9pOvmvbgCd6GlgDGZqddVPtydg8aC9gKDqybaA9EX2nyHWgrhTkxUQ2W9uajL4U0QdAbNsAQFsy3C8wBHTsXQk+BmhpitwE3vgCANSGVi8pfATO207uooh4hqc9vgHxeLzgeXIG+KFIv5Ve7GxujuZ9AwAONYfnVLUHCCHec9vO1vsKALSloq9A7gONhILDm/2yWwYAzETDbeC1wgk7k+/bKCuW4yqwDIxtGBSGzURktNzbdrZeQ8EJIA76CJXCP/kLCpEyUE2VVLWnLRV9Wh7MZPJHPNW3QE2lQ0YgSKpKgLWiV1zftyTCn758XUwZ4lUEdrx+ASW88rx2UJY8AAAAAElFTkSuQmCC&label=Nord&message=v0.2.0&colorA=4c566a&colorB=88c0d0"/>
  </a>
</p>

---

# 0.2.0

_2017-01-09_

## Improvements

❯ Characters under block cursors are now colored darker (`nord1`) while the block cursor is visible to achieve a optimal contrast and to avoid unreadability due to the same cursor- and foreground color (`nord4`). (@svengreb / @kepbod, #2, 8b6e64b0)

<p align="center">
  <strong>Before</strong><br />
  <picture>
    <img src="https://cloud.githubusercontent.com/assets/7836623/21782939/83fc004c-d6b5-11e6-84ac-a05fdee61fa5.png"/><br />
  </picture>
  <strong>After</strong><br />
  <picture>
    <img src="https://cloud.githubusercontent.com/assets/7836623/21782954/942b03aa-d6b5-11e6-9f52-af2a2d737b1a.png"/>
  </picture>
</p>

### Documentation

❯ Added a screenshot to the README [activation](https://github.com/nordtheme/iterm2/blob/develop/README.md#activation) description (@svengreb, 669c307d)

# 0.1.0

_2016-11-22_

## Features

❯ Implemented the main XML (Property List) color scheme file [`Nord.itermcolors`](https://github.com/nordtheme/iterm2/blob/develop/src/xml/Nord.itermcolors). (@svengreb, #1, 80d14907)

![](https://raw.githubusercontent.com/nordtheme/iterm2/develop/src/assets/scrot-colortest.png)
![](https://raw.githubusercontent.com/nordtheme/iterm2/develop/src/assets/scrot-htop.png)

Detailed information about features, supported languages and install instructions can be found in the [README](https://github.com/nordtheme/iterm2/blob/develop/README.md#installation) and in the [project wiki](https://github.com/nordtheme/iterm2/wiki).

# 0.0.0 (2016-11-22)

**Project Initialization**

<p align="center">
  <picture>
    <source srcset="https://raw.githubusercontent.com/nordtheme/assets/main/static/images/elements/separators/iceberg/footer/dark/spaced.png" srcset="https://raw.githubusercontent.com/nordtheme/assets/main/static/images/elements/separators/iceberg/footer/dark/spaced-2x.png 2x" width="100%" media="(prefers-color-scheme: light), (prefers-color-scheme: no-preference)" />
    <source srcset="https://raw.githubusercontent.com/nordtheme/assets/main/static/images/elements/separators/iceberg/footer/light/spaced.png" srcset="https://raw.githubusercontent.com/nordtheme/assets/main/static/images/elements/separators/iceberg/footer/light/spaced-2x.png 2x" width="100%" media="(prefers-color-scheme: dark)" />
    <img src="https://raw.githubusercontent.com/nordtheme/assets/main/static/images/elements/separators/iceberg/footer/dark/spaced.png" srcset="https://raw.githubusercontent.com/nordtheme/assets/main/static/images/elements/separators/iceberg/footer/dark/spaced-2x.png 2x" width="100%" />
  </picture>
</p>

<p align="center">
  Copyright &copy; 2016-present <a href="https://www.svengreb.de" target="_blank">Sven Greb</a>
</p>

<p align="center">
  <a href="https://github.com/nordtheme/iterm2/blob/develop/license" target="_blank">
    <img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/>
  </a>
  <a href="https://www.svengreb.de">
    <img src="https://img.shields.io/static/v1.svg?style=flat-square&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAABmJLR0QA/wD/AP+gvaeTAAABMklEQVQ4jcWQvUoDQRRGz52s5IfVIiDWPkGKFFaCIVaGdIagjcFAwICFb7DvIK6QQlNpY2UQLMQVBbEQ0SewFkGbKCQmOzaTJay7/lR+zTAf9xwuF/47Mv45rdezqWEq72v/RWZnHgqOMwDwHMfSj085JSqb6Pu38we7r18E3nqzhmYbsE11rxKsAvhDfQiSM30XYbOw57YDwfnaRl6U3ABWaMNn806H+oGPzBX3d+4UgChZiYBHYBgGsBLoKoAyhR0x9G20Zmpc4P1ZoMQDcwMNclFrdhBKv6M5WWi7ZQGtjEUn35IV4OwnVjSX/WGmKqCDDUa5rmyle3bvGFiMg3WGUsF1u0EXHoqTRMGRgkAy2eugKZrqijRLYThWANBpNDL2h3UE0J0YLJdbrfe42f/NJ0wqY7/KcXKPAAAAAElFTkSuQmCC&label=lovely%20crafted%20in&message=Germany&colorA=4c566a&colorB=88c0d0"/>
  </a>
</p>
