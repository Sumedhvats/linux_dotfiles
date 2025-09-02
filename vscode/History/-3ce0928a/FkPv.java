import java.util.*;

class Solution {
    public ArrayList<ArrayList<Integer>> countFreq(int[] arr) {
        Map<Integer,Integer> map = new HashMap<>();
        for(int i =0;i<arr.length;i++){
            map.put(arr[i],map.getOrDefault(arr[i],0)+1);
        }
        ArrayList<ArrayList<Integer>>ans= new ArrayList<>();

        int i =0;
        for(Map.Entry<Integer,Integer> entry:map.entrySet()){
            ArrayList<Integer>curr=new ArrayList<>();
            curr.add(entry.getKey());
            curr.add(entry.getValue());
            ans.add(curr);
        }return ans;


    }
}