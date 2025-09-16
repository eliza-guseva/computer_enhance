from setuptools import setup, Extension
from Cython.Build import cythonize

extensions = [
    Extension(
        "cython_sum",
        ["cython_sum.pyx"],
        extra_compile_args=["-mavx2", "-O3"]
    )
]

setup(
    ext_modules = cythonize(extensions)
)
