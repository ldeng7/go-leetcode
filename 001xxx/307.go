func isSolvable(words []string, result string) bool {

class Solution {
public:
    int weight[26],maxJ=0;
    bool used[10];
    bool bt(vector<string>& words,string &result,int i,int j,int curSum,int carry)
    {
        if(j==maxJ&&i==0)
            return true;
        if(i==words.size())
        {
            curSum+=carry,carry=curSum/10,curSum%=10;
            if(used[curSum]&&weight[result[j]-'A']==curSum&&bt(words,result,0,j+1,0,carry))
                    return true;
            else if(!used[curSum]&&weight[result[j]-'A']==-1)
            {
                weight[result[j]-'A']=curSum,used[curSum]=true;
                if(bt(words,result,0,j+1,0,carry))
                    return true;
                used[curSum]=false,weight[result[j]-'A']=-1;
            }
            return false;
        }
        if(j>=words[i].length())
            return bt(words,result,i+1,j,curSum,carry);
        if(weight[words[i][j]-'A']!=-1)
            return bt(words,result,i+1,j,curSum+weight[words[i][j]-'A'],carry);
        else
            for(int val=0;val<=9;val++)
            {
                if(used[val]||val==0&&j==words[i].length()-1)
                    continue;
                weight[words[i][j]-'A']=val,used[val]=true;
                if(bt(words,result,i+1,j,curSum+val,carry))
                    return true;
                weight[words[i][j]-'A']=-1,used[val]=false;
            }
        return false;
    }
    bool isSolvable(vector<string>& words, string result)
    {
        memset(used,false,sizeof used),memset(weight,-1,sizeof weight);
        for(string &w:words)
            maxJ=max(maxJ,(int)w.length()),reverse(w.begin(),w.end());
        if(result.length()>maxJ+1||result.size()<maxJ)
            return false;
        maxJ=max(maxJ,(int)result.length()),reverse(result.begin(),result.end());
        return bt(words,result,0,0,0,0);
    }
};
