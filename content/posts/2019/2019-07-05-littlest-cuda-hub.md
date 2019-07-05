---
title: "Setting up GPU Data Science environments for hackathons"
subtitle: "AKA Littlest Cuda Hub"
date: 2019-07-05T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - rapids
  - iris
  - aws
  - jupyterhub
  - cuda
  - gpu
thumbnail: jupyter
---

## Background

In my first week working at Nvidia I have been spending some time with my previous colleagues at the Met Office to explore how the two organisations can collaborate. 

The theme for the week was to run a hackathon to explore how GPUs could be used to accelerate existing tools and workflows within the Met Office. To begin trying things out we needed a collaborative environment with access to GPUs. For this the Met Office kindly provided a GPU EC2 instance on AWS and we set up RAPIDS and the [Littlest JupyterHub](https://tljh.jupyter.org/en/latest/) (tljh) to expose it to collaborators.

## Setup

Let's walk through the steps to create our GPU enabled Littlest JupyterHub.

### Infrastructure

First we will need to build a ubuntu server with GPUs available. As the Met Office folks have provided access to an AWS account we will go through the excellent instructions from the [tljh docs](https://tljh.jupyter.org/en/latest/install/amazon.html) to set up on AWS with a few adjustments below.

- At step 6 we selected a `p3.8xlarge` instance as it comes with four Nvidia Tesla V100 GPUs. 
- At step 7 we skipped entering the user data (bootstrapping script) as we wanted to run this ourselves manually later.
- At step 8 we selected 100GB to give ourselves reasonable space for data.
- Stop after step 16. As we skipped step 7, steps 17-19 do not apply and we need to do some manual configuration before continuing.

### Drivers

Next you need to ssh to the box with the ssh key you selected during setup. In order to make use of GPU enabled libraries we needed to install the NVidia CUDA drivers.

- Visit https://www.nvidia.co.uk/Download/index.aspx
  - Select the appropriate GPU for the instance (in this case Tesla V100)
  - Select Linux 64-bit Ubuntu 18.04 
    - _WARNING: Do not select the generic linux 64-bit driver as it will not work on ubuntu_
  - Copy the download link and run `wget` on the EC2 instance to download the `.deb` file.
  - Install your `.deb` file on the server (in our case we ran `dpkg -i nvidia-diag-driver-local-repo-ubuntu1804-418.67_1.0-1_amd64.deb` but your driver file name may differ)
- This will have installed a new repository but not the actual drivers so let's also do that.
  - `apt-get update`
  - `apt-get install cuda-drivers`
  - We now need to `reboot` for the drivers to be loaded into the kernel

### JupyterHub

Now that we have drivers in place and have rebooted we can install the Littlest JupyterHub as instructed in the documentation. In step 7 we skipped entering the installation command in the user data field, so let's run that manually now.

```
$ curl https://raw.githubusercontent.com/jupyterhub/the-littlest-jupyterhub/master/bootstrap/bootstrap.py | sudo python3 - --admin <admin-user-name>
```

Select an admin username that you will remember as you will need that the first time you log in to JupyterHub.

We should now be able to return to the tljh documentation and continue from step 20 by visiting the server's IP address in the browser.

Once you finish this I also recommend upgrading the version of Jupyter Lab to the latest by running `sudo -E pip install -U jupyterlab` from within Jupyter Lab.

### RAPIDS

Now that we have a working enviroment we can get on with the interesting bit of installing the packages from RAPIDS. There is a [useful utility](https://rapids.ai/start.html) on the RAPIDS website for constructing your installation command.

You should select:
- Conda
- RAPIDS Stable
- All packages
- Python 3.7
- CUDA 10.0

Then copy the command, prefix it with `sudo -E`, add `cupy` to the end and run it in Jupyter Lab. For me this command looked like this.

```
$ sudo -E conda install -c nvidia -c rapidsai -c numba -c conda-forge -c pytorch -c defaults \
cudf=0.8 cuml=0.8 cugraph=0.8 python=3.7 cudatoolkit=10.0 cupy
```

### Testing

To make sure things work we can create a couple of notebooks and use some of these libraries.

#### cudf

Let's start with cudf. To check this we can visit the documentation and copy the getting started example.

```python
import cudf, io, requests
from io import StringIO

url="https://github.com/plotly/datasets/raw/master/tips.csv"
content = requests.get(url).content.decode('utf-8')

tips_df = cudf.read_csv(StringIO(content))
tips_df['tip_percentage'] = tips_df['tip'] / tips_df['total_bill'] * 100

# display average tip by dining party size
print(tips_df.groupby('size').tip_percentage.mean())
```

If everything works ok you should see the following output.

```
size
1    21.729201548727808
2    16.571919173482897
3    15.215685473711837
4    14.594900639351332
5    14.149548965142023
6    15.622920072028379
Name: tip_percentage, dtype: float64
```

#### cupy

We can also test cupy out by creating an array and playing around with it.

```python

import cupy as cp
x = cp.arange(6).reshape(2, 3).astype('f')
print(x.sum(axis=1))     
```

This should show

```
array([  3.,  12.], dtype=float32)   
```

## Uses

Now we have a great place to experiment with GPU accelerated Python libraries. Here are some links to posts about some of the things that took place at the Met Office hackathon:

  - Using Cupy in Iris (TODO)
  - Alistair (TODO)
  - 

## Future enhancements

- Rotating GPUs
- GPU dashboard