## 内积

又称数量积，点积

α ⋅ β = ∣ α ∣ ∣ β ∣ c o s θ

对加法满足分配律

### 几何意义
向量α在向量β的投影α ’ (带有方向性）与β的长度乘积

- 若α与β 的夹角为锐角，则其内积为正

- 若α与β的夹角为钝角，则其内积为负

- 若α与β 的夹角为直角，则其内积为0
- x1x2+y1y2

```double Dot(Vector A, Vector B){   ```

​    ``` return A.x*B.x + A.y*B.y; ``` 

```}```



## 外积

又称向量积，叉积

α × β = ∣ α ∣ ∣ β ∣ s i n θ 

θ表示向量 α 旋转到向量 β 所经过的夹角

对加法满足分配律

### 几何意义

向量α 与β 所张成的平行四边形的有向面积

- x1y2-x2y1

```double Cross(Vector A, Vector B){ ```

​     ```  return A.x*B.y-A.y*B.x;```   

``` }```



