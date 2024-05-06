#pragma once
#include <atomic>
// #include <stdio.h>  // for debug

template<typename T>
class sharedPtr
{
/*
    API:
    sharedPtr<T>(T* rawPointer);                        // usage: sharedPtr<T>(new T(...));
    sharedPtr<T>(const sharedPtr<T>& rptr);             // usage: sharedPtr<T> sp2(sp1);
    sharedPtr<T>& operator=(const sharedPtr<T>& rptr);  // usage: auto sp2 = sp1;
    
    T* get() const;
    T& operator*() const;
    T* operator->() const;
    long user_count() const;
    

*/


public:
    sharedPtr(T* rawPointer):
        rawPointer_(rawPointer),
        refCount_(new std::atomic<long>(1))
    {  }

    sharedPtr(const sharedPtr<T>& rptr):
        rawPointer_(rptr.rawPointer_),
        refCount_(rptr.refCount_)
    { 
        (*refCount_)++;    // use atomic-val to replace the mutex-usage
    }

    sharedPtr<T>& operator=(const sharedPtr<T>& rptr)
    {
        // process self-point situation
        if(rptr == *this)
            return *this;
        
        release();

        rawPointer_ = rptr.rawPointer_;
        refCount_ = rptr.refCount_;
        /* ------ race-condition ? YES ------ */
        (*refCount_)++;    // use atomic-val to replace the mutex-usage

        return *this;
    }

    ~sharedPtr()
    { 
        release(); 
    }

    T* get() const { return rawPointer_; }
    T* operator->() const { return this->get(); }
    T& operator*() const { return *rawPointer_; }
    long user_count() const { return *refCount_; }

private:

    void release()
    {
        if(refCount_->fetch_sub(1) == 1)
        {
            delete rawPointer_;
            delete refCount_;
        }

    }

    T* rawPointer_;
    std::atomic<long>* refCount_;
};