head = 150;
perLine = 130;
hiBtnHeight = 95;
hiBtnWidth = 670;
addBtnHeight = 64;
addBtnWidth = 144;
touchMoveCount = 2;
touchSleepMills = 100;
waitMills = 1000;
init(0);


local sz = require("sz")
local pos = sz.pos

function findBtnAndTouch(color,btnWidth,btnHeight)
	x, y = findColorInRegionFuzzy(color, 100, 0, head,  width, height); 

	if x<=-1 or y<=-1 then
		return x,y;
	end

	yp1 = y + btnHeight/2;
	yp2 = y + btnHeight;

	hiBtnPointY = yp1;
	x1, y1 = findColorInRegionFuzzy(color, 100, 0, yp1,  width, yp2); 
	if x1<=-1 or y1<= -1 then
		hiBtnPointY = y - btnHeight/2;
	end

	xp1 = x + btnWidth/2;
	xp2 = x + btnWidth;

	hiBtnPointX = xp1;
	x2, y2 = findColorInRegionFuzzy(color, 100, xp1, hiBtnPointY-1,  xp2, hiBtnPointY+1); 
	if x2<=-1 or y2<= -1 then
		hiBtnPointX = x - btnWidth/2;
	end

	touchDown(hiBtnPointX, hiBtnPointY);
	mSleep(touchSleepMills);
    touchUp(hiBtnPointX,hiBtnPointY);
    return hiBtnPointX, hiBtnPointY;
end

toast("正在打开微信附近的人"); 
r = runApp("com.tencent.mm","com.tencent.mm.plugin.nearby.ui.NearbyFriendsUI");
if r==0 then
	width,height = getScreenSize(); 
	length = math.floor((height-head)/perLine);
    mSleep(5*waitMills);
    toast("开始打招呼"); 
    for i=1,touchMoveCount,1 do
    	touchX = width/2;
	    for i=0,length-1,1 do
	    	mSleep(waitMills);
	        touchY = (head+perLine/2)+ (perLine*i);
	    	touchDown(touchX, touchY);
	    	mSleep(touchSleepMills);
	   		touchUp(touchX, touchY);
		    mSleep(waitMills);
			x,y = findBtnAndTouch(0x45c01a,hiBtnWidth,hiBtnHeight)
		    if x>-1 and y>-1 then
		    	mSleep(waitMills);
		    	x1,y2 = findBtnAndTouch(0x45c01a,addBtnWidth,addBtnHeight)
		    	os.execute("input keyevent 4");
		    	mSleep(waitMills);
		    	os.execute("input keyevent 4");
		    end
		    mSleep(waitMills);
	    	os.execute("input keyevent 4");
	    end
	    mSleep(waitMills);

		local p0 = pos(touchX, (head+ perLine*length))
		local p1 = pos(touchX,head)
	    p0:touchMoveTo(p1,5,touchSleepMills,10);
	    --[[ouchDown(touchX, (head+ perLine*length));
	    mSleep(touchSleepMills);
	    touchMove(touchX,head);
	    mSleep(touchSleepMills);
	    touchUp(touchX,head);
	    ]]
	    mSleep(waitMills);
	end
	toast("打招呼结束"); 
    dialog(x.."  "..y);
else
	dialog("启动失败");
end
