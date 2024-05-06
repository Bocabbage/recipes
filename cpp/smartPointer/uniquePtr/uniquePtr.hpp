#pragma once
#include <utility>

namespace demosp {

template<typename T>
class uniquePtr
{
private:
    T* _ptr;

public:
    // ------- Constructor/Destructor -------
    uniquePtr() = delete;
    uniquePtr(const uniquePtr<T>&) = delete;
    uniquePtr<T>& operator=(const uniquePtr<T>&) = delete;

    uniquePtr(T* raw_ptr):
    _ptr(raw_ptr) 
    { }

    uniquePtr(uniquePtr<T>&& uptr):
    _ptr(uptr.release())
    { }

    uniquePtr<T>& operator=(uniquePtr<T>&& uptr)
    {
        if(this == &uptr) // [TODO] check if correct
            return *this;
        if(_ptr)
            delete _ptr;
        
        _ptr = uptr.release();
        return *this;
    }

    ~uniquePtr() 
    {
        if(_ptr)
            delete _ptr;
    }

    // -------- Methods -------
    T* get() const
    { return _ptr; }

    T* operator->() const 
    { return this->get(); }

    T& operator*() const 
    { return *_ptr; }

    T* release()
    {
        T* tmp = nullptr;
        std::swap(_ptr, tmp);
        return tmp;
    }

    void reset(T* raw_ptr)
    {
        if(_ptr == raw_ptr)
            return;
        auto old_ptr = this->release();
        if(old_ptr)
            delete old_ptr;
        _ptr = raw_ptr;
    }

    void swap(uniquePtr<T>& uptr)
    {
        std::swap(_ptr, uptr._ptr);
    }

};

template<typename T, class... Args>
uniquePtr<T> make_unique(Args&&... args) 
{
    return uniquePtr<T>(new T(std::forward<Args>(args)...));
}

}

