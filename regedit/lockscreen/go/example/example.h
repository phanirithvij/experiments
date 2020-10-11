#ifndef _EXAMPLE_H_
#define _EXAMPLE_H_

#ifndef DLLAPI
#define DLLAPI extern "C" __declspec(dllexport)
#endif

DLLAPI bool example_func();

#endif // _EXAMPLE_H_
