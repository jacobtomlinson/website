---
title: "Sometimes I regret using CalVer"
date: 2023-01-16T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - software development
  - open source
  - python
  - semver
  - calver
---

Over the last few years, many open-source Python projects that I work on have switched to [CalVer](https://calver.org/). I've felt some pain around this, particularly in [Dask](https://dask.org/) and its subprojects. I want to unpack some of my thoughts and feelings around this trend.

```info
Semantic Versioning or [SemVer](https://semver.org/) is where versions increments are calculated based on what has changed in the release, whether it is a major/breaking change, a minor/feature change or just a bug fix, e.g `MAJOR.MINOR.PATCH`.

Calendar Versioning or [CalVer](https://calver.org/) is where a project's versioning scheme simply uses the current date rather than any semantic meaning, e.g `YEAR.MONTH.DAY`.

Some projects use a hybrid of the two. In Dask we use `YEAR.MONTH.PATCH` as it allows us to release multiple times within one day if we need to.
```

I'm going to break this post down into the various takeaways I have from working with CalVer for a while.

## Calendar versions can (and should) still have some semantics

One of the biggest examples folks hold up as a successful CalVer project is [Ubuntu](https://releases.ubuntu.com/). Ubuntu follows the `YY.MM.P` scheme, which looks pretty similar to Dask's `YYYY.MM.I` scheme. However Ubuntu has a bunch of process and behaviors around their versioning that means it does include semantics.

My main regret about how we do things in Dask is that the versions have no semantics at all. Releases happen every two weeks and the version is calculated based on the date of that release. This means the difference between `2022.08.1` and `2022.09.2` is impossible to tell. Was `2022.08.1` the second release in August, or a hotfix for a break in `2022.08.0`? Which release is more stable? Should I expect them to be compatible? There just isn't any information here to make a decision.

Ubuntu releases every 6 months in either April or October. This is fixed and they never deviate, even if reality doesn't quite line up. Ubuntu releases have been known to be late, the spring release might end up dropping in May or June, but the version is still `YY.04.P` because that is what is expected by the community.

Having 6 months between releases means the patch can increment a reasonable number of times and folks can trust that `22.04.8` will be compatible with `22.04.2` and should be more stable.

Ubuntu also makes no guarantees about stability between the 6 monthly releases. This effectively makes the `YY.MM` portion of the version a `MAJOR` semantic with a little date info thrown in. The last portion of the version is explicitly a patch and shouldn't cause breakages or significant changes. Ubuntu could also use `MAJOR.PATCH` in exactly the same way.

Ubuntu also goes out of their way to backport fixes to older versions and signals that certain releases will be supported more strongly with the `LTS` or Long Term Support modifier. This leads many users who value stability to upgrade every two years when a new LTS release drops. This gives them increased confidence in a particular release and also allows them to schedule in the migration overhead well in advance.

## Distributions should be calendar versioned

Ubuntu is also not a library or API, it is a Linux distribution. Distributions are collections of tens, hundreds or thousands of pieces of software that have been curated and tested together. The semantics of major/breaking changes and minor/feature changes are less useful there. Every release will have many changes of all types from many projects in their collection.

Distributions are also highly opinionated and at the same time totally arbitrary. A person or group of people have selected their preferred tools to fulfill the needs of a modern operating system and then have invested time and effort to verify that those tools are compatible. The decisions to add, upgrade or replace various pieces of software are too complex to represent in a versioning scheme. Snapshotting a point in time and saying "this collection is good, we've tested it and are happy" is the best you're going to get. Then following up with patch releases is effectively like saying "remember that good collection we had, we've fixed a few issues with it".

Many projects do the same kind of thing, but at different levels of the software stack. Some curate container images containing software environments, such as the [Pangeo containers](https://github.com/pangeo-data/pangeo-docker-images). Other projects such as [Kubeflow](https://github.com/kubeflow/kubeflow) build platforms and toolkits out of many other pieces of software to serve a particular purpose. These are also distributions and calendar versioning is probably the right choice for them.

One of the reasons behind Dask switching to CalVer was that it provides alternative implementations of the Pandas, NumPy and SciKit-Learn APIs which we refer to as Dask collections. This sort of makes Dask a distribution given that it is vendoring in multiple APIs from other projects. However in hindsight I'd argue that while the Dask collections are a vital part of Dask they are only a small part of the versioned ecosystem. Other components such as the Dask scheduler or deployment tooling should probably have stuck with SemVer given that they are just providing APIs to users and other projects.

## Libraries and APIs should be semantically versioned

As a consumer of many APIs the thing I appreciate most is stability and predictability. It is always frustrating when a new release of an API happens and I have to make changes to my code. The stability of an API is unrelated to the scheme used to version it, but with SemVer I can make an educated guess about how much time it should take me to upgrade. This could be a few minutes to bump a patch release or a few weeks to refactor things to use a new major release.

A common argument I see from folks in the PyData space against SemVer is that it isn't done well in our community and so we shouldn't bother. From time to time a minor or patch release will happen that breaks things, but in my opinion that's ok, mistakes happen. As long as the project issues a follow up patch release to revert the breakage it's not a huge issue. The important thing is that you are trying to stick to the semantics, and if a release does something unexpected it is a broken release.

By releasing software that uses SemVer you are signaling to your community that you have some constraints. Depending on the types of changes you make, you will signal that to them with the version number you release. If you regularly make major version bumps with breaking changes the community will not be happy about this and likely push back.

In my opinion, CalVer signals to your community that anything could happen at any time and that you have no interest in the effect that has on your users. The project has been tested to be working on a given date, but it is an exercise for the user to figure out how much effort it would take for them to upgrade. I don't feel this is very respectful of users time and effort.

## Version schemes should be SemVer compliant

While I agree there is a time and place for CalVer there are many projects and ecosystems that are tightly wedded to SemVer. CalVer and SemVer can be compatible, but with Dask we fell into a gotcha.

SemVer doesn't allow for zero padding of versions. However, it might be tempting to zero pad the year or month in your version so that versions sort lexicographically. In Dask we went with `YYYY.MM.P`, so the month gets zero padded for most of the year. We later found that [helm](https://helm.sh/) enforces SemVer complaint versions in its chart repos, so when we automated our chart publishing to push out a new image with each Dask release we got errors saying the Dask release version was not valid.

The quick fix for this is to use `YYYY.M.P` for our Helm Chart releases. However this introduces an inconsistency that frustrates me because we need to remember this special case and work around it indefinitely.

It's also interesting to note that PyPI doesn't allow zero padding as part of complying with [PEP440](https://peps.python.org/pep-0440/). But it automatically strips the zero if you try and push a release with a passes version.

## CalVer doesn't make releases more regular

The driving motivation for Conda to switch to CalVer in [CEP 8](https://github.com/conda-incubator/ceps/blob/main/cep-8.md) was to "remove ambiguity/maintainer guesswork of when and what warrants a release". Much of that CEP goes on to desribe how releases should be created bi-monthly and doesn't actually mention any pros/cons of removing semantics from the versioning.

Any project can switch from an ad-hoc release cycle to a periodic release cycle. You don't need to use CalVer to do this. I think it is easy to be inspired by Ubuntu's versioning and release cycle and adopt the same thing, but I think semantics and release candence should be considered as two separate subjects.

Dask behaves in a similar way just a little faster, the core projects are released every two weeks regardless of what has been done in that time. Sometimes the release gets delayed if a maintainer demonstrates the the current trunk is not stable, and sometimes release happen in between to fix bugs. But due to using CalVer it isn't transparent to the user community what implications each release has for them.

```info
I find [Hypothesis](https://github.com/HypothesisWorks/hypothesis) an interesting project because they [release on every single stable commit](https://github.com/HypothesisWorks/hypothesis/releases). This is another effort to remove the ambiguity around when a maintainer should release. They use SemVer so that it is clear what has changed between the releases.
```

## CalVer versions are not notable

Dask made 27 releases during 2022, but that's all I know about them.

A common reason to upgrade is when you find a bug and are on an old release. The bug has potentially been fixed in a newer release and I should verify before opening an issue. But if I were a user and I was still using a version from 2021 where do I upgrade to?

Maybe I jump to the latest version and find that a bunch of stuff has changed. Do I fix up my code to take the breakages into account? Or do I jump somewhere in between to see if my code still works and the bug has been fixed? Maybe I choose to jump to the release in between where I am and the latest? Does that fix the bug without introducing breaks? If not, do I jump again? Did I just reinvent bisect but for releases?

You get the idea.

If the project used SemVer the user could just move to the latest release for the major version they were on and try it out. If it works then great! If not they know they have to do some migration work to a newer major version.

As a maintainer I think I would find it easier to remember what had caused the major bump in a Semver project too. So if I was providing user support I could maybe even tell them if their code would work with a newer release. With CalVer I find it very hard to remember what happened between releases in the past without going back through the entire changelog.

## Switching back is hard

Moving to CalVer can be pretty easy. Given that many CalVer schemes are valid SemVer schemes and most project's major version is lower than the current year you can just put out a CalVer release. That's it.

Dask's final SemVer release was `2.30.1` and the first CalVer release was `2020.12.0`. From a SemVer perspective Dask incremented the `MAJOR` version by two thousand and eighteen and then the `MINOR` version by twelve. This means sorting works correctly and package managers happily install the CalVer release as the most current version. Then further CalVer releases follow the same rules of sorting by changes to the most significant number, the same way dates work.

However going back isn't as straight forward. If hypothetically Dask chose to switch back SemVer then the next semantic version would be `3.0.0` (or maybe `4.0.0` if you count the CalVer releases as v3). But that wouldn't be sorted correctly and would appear below all of the CalVer releases.

The only options are to jump forward to an even higher major version that is clearly not a date or to retcon the CalVer releases and yank/republish existing CalVer releases with a new SemVer name. Neither of which are good options.


## CalVer can imply semantics where there are none

After originally publishing this article a reader reached out to me and made the following comment:

> Speaking now as a library consumer: I'm not a Dask, nor a k8s expert, but were I to try dask-kubernetes, currently at 2022.12.0, I'd feel confident targeting Dask and k8s of similar December 2022 vintage.

All projects using CalVer use a common set of dates which can lead users to assume projects with near dates are compatible, but this assumption worries me. Sure they got released at the same time, but `dask-kubernetes` specifies which `dask` versions it works within its requirements. Compatible versions of each should be installed via the pip/conda solver. It shouldn't be down to the user to make this judgement call.

Sure all December 2022 releases were tested and authored around the same time, but it doesn't mean the latest versions of things were tested. Python `3.11` was released in October 2022 but dask's `2022.10.2` release did not work with Python `3.11`. It wasn't until `2022.12.1` that `3.11` support was added. If Python used CalVer would users assume the same immediate compatibility?

I've also been in a position where `dask` is on a 2023 release already but another subproject like `dask-cloudprovider` is still a few months behind with a `2022` release. Does that mean `dask-cloudprovider` is out of date or won't work with current `dask`? No, of course not. It just means different projects have different complexities and some need changing more than others. Having many projects all using CalVer causes folks to take meaning from this and assume that similar dates are compatible. But no such promise is being made by the maintainer.

Yet I've had users reach out to me asking me to make a release of projects that have had no changes just so the dates line up.

## Wrap up

For any distribution-like project CalVer is a sensible way to version your releases. But there is much more to releasing software than which scheme you use. Building up trust with your community by being transparent, consistent and punctual is crucial.

For me I am still on the fence about whether Dask's overall move to CalVer was a bad choice. But for the majority of Dask subprojects I maintain like [dask-kubernetes](https://kubernetes.dask.org/) and [dask-cloudprovider](https://cloudprovider.dask.org/en/latest/) I regret the switch.
