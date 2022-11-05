import "List.hs";
import "../../../libs/SortType.hs";

data ListPair = ListPair (List SortType) (List SortType);

main :: IO ();
main = do {
    l <- return (getList 128);
    l <- return (quickSort l);

    printList l;
    putStr "\n";
};

quickSort :: List SortType -> List SortType;
quickSort l = case l of {
    Nil -> Nil;
    Cons x xs -> case divideList x xs of {
        ListPair left right -> linkLists (quickSort left) (Cons x (quickSort right));
    };
};

divideList :: SortType -> List SortType -> ListPair;
divideList mid l = foldl (getElToLists mid) (ListPair Nil Nil) l;

getElToLists :: SortType -> ListPair -> SortType -> ListPair;
getElToLists mid pair x = case pair of {
    ListPair left right -> case cmp x mid of {
        True -> ListPair (Cons x left) right;
        False -> ListPair left (Cons x right);
    };
};


