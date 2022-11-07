data SortType = SortType Int;

cmp :: SortType -> SortType -> Bool;
cmp x1 x2 = (getNum x1) < (getNum x2);

getList :: Int -> List SortType;
getList n = reverseList (intListToElList (nNumbers n));

intListToElList :: List Int -> List SortType;
intListToElList = map getEl;

reverseList :: List SortType -> List SortType;
reverseList l = foldl append Nil l;

linkLists :: List SortType -> List SortType -> List SortType;
linkLists left right = case left of {
    Nil -> right;
    Cons x xs -> Cons x (linkLists xs right);
};

append :: List SortType -> SortType -> List SortType;
append l x = Cons x l;

printList :: List SortType -> IO ();
printList l = case l of {
    Nil -> putStr "";
    Cons x xs -> do {
        printEl x;
        case xs of {
            Nil -> putStr "";
            Cons y ys -> do {
                putStr ", ";
                printList xs;
            };
        };
    };
};

printEl :: SortType -> IO ();
printEl x = putInt (getNum x);

getNum :: SortType -> Int;
getNum x = case x of {
    SortType n -> n;
};

getEl :: Int -> SortType;
getEl n = SortType n;

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

numbers :: List Int;
numbers = numbersFrom 0;

nNumbers :: Int -> List Int;
nNumbers n = take n numbers;
