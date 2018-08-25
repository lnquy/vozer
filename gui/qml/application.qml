import QtQuick 2.11
import QtQuick.Controls 2.3
import QtQuick.Controls.Material 2.3
import QtQuick.Controls.Universal 2.3
import QtQuick.Layouts 1.11

ApplicationWindow {
    visible: true
    title: "vozer 0.1.0 - https://github.com/lnquy/vozer"
    property int margin: 11
    minimumWidth: 640
    minimumHeight: 345
    height: 345

    ColumnLayout {
        id: mainLayout
        transformOrigin: Item.Center
        anchors.fill: parent
        anchors.margins: margin

        GroupBox {
            id: gbCrawl
            y: 0
            height: 115
            transformOrigin: Item.Top
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            Layout.rowSpan: 1
            title: ""
            Layout.fillWidth: true

            ColumnLayout {
                id: columnLayout
                spacing: 10
                anchors.fill: parent

                RowLayout {
                    id: rowCrawl
                    anchors.right: parent.right
                    anchors.left: parent.left
                    anchors.top: parent.top
                    TextField {
                        id: tfURL
                        placeholderText: "Link to VOZ thread"
                        Layout.fillWidth: true
                    }
                    Button {
                        id: btnCrawl
                        text: "Crawl"
                    }
                }

                RowLayout {
                    id: rowArgs
                    anchors.top: rowCrawl.bottom
                    anchors.topMargin: 10
                    anchors.right: parent.right
                    anchors.bottom: parent.bottom
                    anchors.left: parent.left
                    CheckBox {
                        id: cbCrawlImages
                        text: qsTr("Crawl images")
                        leftPadding: 0
                        topPadding: 6
                        checked: true
                    }

                    CheckBox {
                        id: cbCrawlUrls
                        text: qsTr("Crawl URLs")
                        checked: true
                    }

                    Label {
                        id: lbOutput
                        text: qsTr("Save to")
                        rightPadding: 0
                        leftPadding: 10
                    }

                    TextField {
                        id: tfURL1
                        placeholderText: "Path to output directory"
                        Layout.fillWidth: true
                    }
                }
            }

        }

        GroupBox {
            id: gbAdvOptions
            y: 125
            Layout.fillHeight: true
            transformOrigin: Item.Top
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            leftPadding: 0
            topPadding: 41
            title: "Advanced options"
            Layout.fillWidth: true

            ColumnLayout {
                id: columnLayout1
                anchors.leftMargin: 10
                anchors.fill: parent
                spacing: 10

                RowLayout {
                    id: rowCrawlByPages

                    Label { id: lbCrawlByPages; text: "Crawl by pages list";rightPadding: 10 }
                    TextField { id: tfCrawlByPages; leftPadding: 10; Layout.fillWidth: true; placeholderText: "List of page numbers" }
                }

                RowLayout {
                    id: rowCrawlByRange
                    Label { id: lbCrawlByRange; text: "Or crawl by range";rightPadding: 18 }
                    TextField { id: tfCrawlByRangeFrom ;width: 155; text: ""; Layout.fillWidth: true;placeholderText: "From page" }

                    Label {
                        text: qsTr("-")
                    }

                    TextField {
                        id: tfCrawlByRangeTo
                        width: 155
                        Layout.fillWidth: true
                        transformOrigin: Item.Center
                        placeholderText: "To page"
                    }

                }

                RowLayout {
                    id: rowSeparator
                    height: 5
                    spacing: 0
                    Layout.fillHeight: false
                    Layout.fillWidth: true

                    Rectangle {
                        id: rectangle
                        width: 200
                        height: 1
                        color: "#cccccc"
                        border.color: "#cccccc"
                        Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                        Layout.fillWidth: true
                    }


                }

                RowLayout {
                    id: rowCrawlByRange1
                    Label {
                        id: lbCrawlByRange1
                        text: "Workers"
                        rightPadding: 0
                    }

                    TextField {
                        id: tfCrawlByRangeFrom1
                        width: 155
                        text: "10"
                        placeholderText: "# of workers to crawl data"
                        Layout.fillWidth: true
                    }

                    Label {
                        text: qsTr("Retries")
                        rightPadding: 0
                        leftPadding: 15
                    }

                    TextField {
                        id: tfCrawlByRangeTo1
                        width: 155
                        text: "20"
                        placeholderText: "# of times to retry when failed"
                        transformOrigin: Item.Center
                        Layout.fillWidth: true
                    }
                }
            }
        }
    }
}
