package code05

import "sync"

type ReadWriteMutex struct {
    readersCounter int
    readerLock     sync.Mutex
    globalLock     sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
    rw.readerLock.Lock()
    rw.readersCounter++
    if rw.readersCounter == 1 {
        rw.globalLock.Lock()
    }
    rw.readerLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
    rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
    rw.readerLock.Lock()
    rw.readersCounter--
    if rw.readersCounter == 0 {
        rw.globalLock.Unlock()
    }
    rw.readerLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
    rw.globalLock.Unlock()
}
