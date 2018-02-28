# FilterHacker

This project was a submission to CruzHacks 2018. Visit it on
[Devpost](https://devpost.com/software/unnamed-photo-thing-wip).

Filter Hacker is a photo sharing social network prototype. It is similar to a
service like Instagram, except instead of using predefined filters, users can
write their own code to make any graphical transformation they want to their
photo. The name is in reference to users "hacking" their own custom filters.

This project is intended to be fun, of course, but it also aims to connect
coding to a less explored medium. Many young people that spend most of their
screen time on photo sharing apps on smart devices may be missing opportunities
to see what writing their own software or code is like, and this project aims to
bridge that gap.

This project was difficult to write (especially in less than two days) because
accepting code from users is such an enormous security risk. The potentially
unsafe parts (nearly everything) are in [the unsafe
repo](https://github.com/spencer-p/unsafe.FilterHacker), intended to be served
from a different subdomain. This complexity makes me unlikely to visit this
project again in its current form, but it may resurface as a mobile app (where
JS tricks can't escape a sandboxed webview).
