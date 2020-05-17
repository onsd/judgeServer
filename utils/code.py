n,k=map(int,input().split())
a=list(map(int,input().split()))
cnt=k
for i in range(n):
    cnt +=a[i]
    if cnt<=0:
        print('No')
        exit()
print('Yes'if cnt>k else 'No')