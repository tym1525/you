
|ID|int64|订单编号||-------|
|Status|int32|订单状态|"1：挂单2：完全成交3：拒绝4：已撤单/完全撤单5：系统暂时保留指令6：部分成交7：已替换(修改订单时会先撤单然后重新挂单，被撤掉的订单状态就是已替换)8：正在处理9：已过期/无效10：正在取消11：正在替换"|-------|
|StatusDesc|string|委托状态描述||-------|
|Side|bool|买卖||-------|
|OpenClose|int32|开平||-------|
|Quantity|int32|下单数量||-------|
|OrderType|int32|订单类型||-------|
|OrderFlag|int32|有效标记||-------|
|DealedQty|int32|成交数量||-------|
|DealedPrice|float64|成交价格||-------|
|DisplayQty|int32|定义的挂单显示数量||-------|
|SecurityDesc|string|合约名称||-------|
|ErrorCode|int32|错误编号|如果=0，没有错误信息|-------|
|ErrorSource|string|错误来源||-------|
|ErrorCodeDesc|string|错误编号描述||-------|
|CreateTimeDisp|string|订单创建时间||-------|
|ProcessTimeDisp|string|订单被处理时间||-------|
|ContractMsg|*CtrMsg|合约信息|点击链接了解详情|-------|


