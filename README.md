#### Contract Match

I document my very late solution for challenge in one of coding assessment platform that I did yesterday.

so the requirement is to count how many character that match with the contract given.

##### requirement: 
    
    -   "0" is stand for vowel letters which contain a, e, i, o, u
        and "1" is stand for consonant.
    
    -   if contract given is "001" then we need to find how many chacracter that matched in the raw data.

Example:
    contract:  ```00```,
    rawData: ```aabbuu```,
    expected: ```2```

why ? cz the rawData with ```aa``` and ```uu``` are meet the contract with ```00 = vowel vowel```
while others like ```ab, bb, bu``` are did't meet the contract.

Finally, I hope you understand what I want to share here, 
I believe there're another ways that probably much better than this solution.

so feel free to add some feedback and happy coding ☕

cheers!
