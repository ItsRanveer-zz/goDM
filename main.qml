import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Window 2.1

Window {
	id: window1
	width: 1000
    height: 500
	x: 200
	y: 200
	
	Rectangle {
		id: rectangle1
	    width: 1000
	    height: 500
		
	    GridView {
	        id: gridView1
	        x: 31
	        y: 63
	        width: 939
	        height: 374
	        clip: false
	        visible: true
	        highlightFollowsCurrentItem: true
	        cacheBuffer: 323
	        transformOrigin: Item.Center
	        cellHeight: 70
	
	        Text {
	            id: text1
	            x: 43
	            y: 71
	            width: 111
	            height: 35
	            text: qsTr("Enter URL")
	            font.pixelSize: 23
	        }
	
	        TextField {
	            id: textField1
	            x: 196
	            y: 71
	            width: 547
	            height: 35
	            placeholderText: qsTr("Enter URL")
	        }
	
	        ProgressBar {
	            id: progressBar1
	            x: 196
	            y: 147
	            width: 547
	            height: 33
	        }
	
	        Text {
	            id: text2
	            x: 425
	            y: 147
	            width: 89
	            height: 33
	            verticalAlignment: Text.AlignVCenter
	            horizontalAlignment: Text.AlignHCenter
	            textFormat: Text.AutoText
	            z: 1
	            font.pixelSize: 14
	        }
	
	        Button {
	            id: button1
	            x: 785
	            y: 71
	            width: 113
	            height: 35
	            text: qsTr("Download")
				enabled: true
	            onClicked: download.startDownload(textField1, progressBar1, text2, button1, button2, button3, button4)
	        }
	
	        Button {
	            id: button2
	            x: 196
	            y: 269
	            width: 146
	            height: 35
	            text: qsTr("Pause")
				enabled: false
				onClicked: download.buttonClicked(0, 0, 0, 0)
	        }
	
	        Button {
	            id: button3
	            x: 396
	            y: 269
	            width: 146
	            height: 35
	            text: qsTr("Resume")
				enabled: false
				onClicked: download.buttonClicked(0, 0, 0, 0)
	        }
	
	        Button {
	            id: button4
	            x: 597
	            y: 269
	            width: 146
	            height: 35
	            text: qsTr("Cancel")
				enabled: false
				onClicked: download.buttonClicked(1, 0, 0, 0)
	        }
	
			Image {
	            id: image
	            x: 785
	            y: 211
	            z: -1
	            source: "gopher.jpg"
	        }
			
	    }
	
	}
	
}