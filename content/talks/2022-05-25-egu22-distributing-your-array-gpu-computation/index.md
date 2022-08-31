---
title: "GPU accelerating your computation in Python"
date: 2022-05-25T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: European Geosciences Union (EGU) General Assembly 2022
  link: https://egu22.eu/
  type: Talk
  location: Online
length: 5
abstract: true
slides: https://speakerdeck.com/jacobtomlinson/gpu-accelerating-your-computation-in-python
---

## Talk abstract

There are many powerful libraries in the Python ecosystem for accelerating the computation of large arrays with GPUs. We have CuPy for GPU array computation, Dask for distributed computation, cuML for machine learning, Pytorch for deep learning and more. We will dig into how these libraries can be used together to accelerate geoscience workflows and how we are working with projects like Xarray to integrate these libraries with domain-specific tooling. Sgkit is already providing this for the field of genetics and we are excited to be working with community groups like Pangeo to bring this kind of tooling to the geosciences.

**How to cite:** Tomlinson, J.: Distributing your GPU array computation in Python, EGU General Assembly 2022, Vienna, Austria, 23–27 May 2022, EGU22-7610, https://doi.org/10.5194/egusphere-egu22-7610, 2022.

## Session abstract

Current pre-exascale computing systems, and the strong push towards exascale warrant substantial efforts to improve the geoscientific software infrastructure used for Earth System Model (ESM) development, data analysis, and storage. The Exascale era opens a range of opportunities, including increased domain size, simulation duration, model resolution, large ensembles, and new physics. This session will discuss challenges and solutions involving domain scientists, applied mathematicians, computer scientists, HPC, and compression experts.
Contributions address challenges and advances to achieve exascale-readiness geoscience disciplines, methods, and technologies.

Pangeo (pangeo.io) is a community of researchers and developers that tackle these issues in a collaborative manner using a growing Python ecosystem whose core tools include xarray, Iris, DASK, Jupyter, Zarr and INTAKE. Many contributors to this session will share novel tools within the Pangeo ecosystem devoted to Atmosphere, Ocean and Land Models, Satellite Observations, HPC, Cloud computing, Machine Learning, and Scalable scientific computing.

This session also considers how geoscientists can shift towards greener computing by adopting modern data compression techniques including, though not limited to: algorithmic advances, assessments of data storage sustainability, compression efficiency and speed in software and/or hardware, interoperability issues, remote sensing applications, and support in widely used languages (e.g., C/C++, Fortran, Java, Python), data storage formats (e.g., HDF, netCDF, Zarr), and Open Source workflows (e.g., CDO, NCO, Pangeo, Ruby, Xarray).

All authors in this session have the option to submit Jupyter notebooks of their work; the best five will be selected as part of the Pangeo applications gallery of EGU22. Examples of previous galleries are at http://gallery.pangeo.io.
