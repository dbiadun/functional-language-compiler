return :: a -> IO a;
putStr :: String -> IO ();
putInt :: Int -> IO ();
getLine :: IO String;

-- Go functions
tinySleep :: Int -> IO ();

-- TinyGo microcontroller functions
tinyLED :: IO Int;
tinyLED2 :: IO Int;
tinyBUTTON :: IO Int;

tinyConfigure :: Int -> Int -> IO ();
tinyPinInput :: IO Int;
tinyPinInputPullup :: IO Int;
tinyPinInputPulldown :: IO Int;
tinyPinOutput :: IO Int;

tinyPinRising :: IO Int;
tinyPinFalling :: IO Int;
tinyPinToggle :: IO Int;

tinyGet :: Int -> IO Bool;
tinyLow :: Int -> IO ();
tinyHigh :: Int -> IO ();
tinySetInterrupt :: Int -> Int -> (Int -> IO ()) -> IO ();

tinySetTimer :: Int -> Int -> IO () -> IO ();
tinyStopTimer :: Int -> IO ();
tinyStartTimer :: Int -> IO ();
