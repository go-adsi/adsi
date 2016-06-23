package adsi

type action func() error

func run(a action) error {
	/*
		// COM is initialized per-thread so we want to prevent the scheduler from
		// switching us over to another thread while we're executing.
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		// TODO: Determine whether the cost of using a dedicated thread is less than
		// the cost of setting up and tearing down the COM interface for every call.

		// FIXME: Find out whether CoInitializeEx can handle simultaneous calls. If
		//        not, guard it with a global mutex.

		// Initialize COM with a multithreaded compartment.
		// See: https://msdn.microsoft.com/en-us/library/ms809971
		if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED|ole.COINIT_DISABLE_OLE1DDE); err != nil {
			oleerr := err.(*ole.OleError)
			// S_FALSE           = 0x00000001 // CoInitializeEx was already called on this thread
			if oleerr.Code() != ole.S_OK && oleerr.Code() != 0x00000001 {
				return err
			}
		}

		// From [MSDN](https://msdn.microsoft.com/en-us/library/windows/desktop/ms688715):
		//
		// Closes the COM library on the current thread, unloads all DLLs loaded by
		// the thread, frees any other resources that the thread maintains, and
		// forces all RPC connections on the thread to close.
		//
		// A thread must call CoUninitialize once for each successful call it has
		// made to the CoInitialize or CoInitializeEx function, including any call
		// that returns S_FALSE.
		defer ole.CoUninitialize()
	*/

	return a()
}
