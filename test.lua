head = 145;
perLine = 128;
hiBtnHeight = 95;
hiBtnWidth = 670;
addBtnHeight = 64;
addBtnWidth = 144;
init(0);

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
	mSleep(30);
    touchUp(hiBtnPointX,hiBtnPointY);
    return hiBtnPointX, hiBtnPointY;
end

r = runApp("com.tencent.mm","com.tencent.mm.plugin.nearby.ui.NearbyFriendsUI");
if r==0 then
	width,height = getScreenSize(); 
	length = (height-head)/perLine
    mSleep(3000);
    for i=1,length,1 do
    	mSleep(1000);
    	touchDown(width/2, (head+(length/2))*i);
    	mSleep(30);
   		touchUp(width/2, (head+(length/2))*i);
	    mSleep(3000);
		x,y = findBtnAndTouch(0x45c01a,hiBtnWidth,hiBtnHeight)
	    if x>-1 and y>-1 then
	    	mSleep(3000);
	    	x1,y2 = findBtnAndTouch(0x45c01a,addBtnWidth,addBtnHeight)
	    	os.execute("input keyevent 4");
	    	mSleep(1000);
	    	os.execute("input keyevent 4");
	    end
	    mSleep(2000);
    	os.execute("input keyevent 4");
    end
	
    dialog(x.."  "..y);
else
	dialog("启动失败");
end
