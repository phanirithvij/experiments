import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

void main() {
  SystemChrome.setEnabledSystemUIOverlays([]);
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Experiments',
      theme: ThemeData.dark(),
      home: MyHomePage(title: 'FlutterExps'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List<PageControllerC> _controllers;
  PageController _rowController;
  ValueNotifier<int> currIdxNotifier = ValueNotifier(0);

  @override
  void initState() {
    _controllers = [
      PageControllerC(
        controller: PageController(keepPage: true),
        recorded: 0,
      ),
      PageControllerC(
        controller: PageController(keepPage: true),
        recorded: 1,
      ),
      PageControllerC(
        controller: PageController(keepPage: true),
        recorded: 2,
      ),
    ];
    _rowController = PageController();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: PageView(
        controller: _rowController,
        onPageChanged: (pno) {
          setState(() {
            currIdxNotifier.value = pno;
            print("${currIdxNotifier.value} horizz");
          });
        },
        children: [
          ColPageView(
            idx: 0,
            notifier: currIdxNotifier,
            controllers: _controllers,
            children: <Widget>[
              ColoredWidget(
                color: Colors.orange[50],
                direction: "0,0",
              ),
              ColoredWidget(
                color: Colors.orange[100],
                direction: "0,1",
              ),
              ColoredWidget(
                color: Colors.orange[200],
                direction: "0,2",
              ),
              ColoredWidget(
                color: Colors.orange[300],
                direction: "0,3",
              ),
            ],
          ),
          ColPageView(
            notifier: currIdxNotifier,
            idx: 1,
            controllers: _controllers,
            children: [
              ColoredWidget(
                color: Colors.green[100],
                direction: "1,0",
              ),
              ColoredWidget(
                color: Colors.green[200],
                direction: "1,1",
              ),
              ColoredWidget(
                color: Colors.green[300],
                direction: "1,2",
              ),
            ],
          ),
          ColPageView(
            notifier: currIdxNotifier,
            idx: 2,
            controllers: _controllers,
            children: [
              ColoredWidget(
                color: Colors.teal[100],
                direction: "2,0",
              ),
              ColoredWidget(
                color: Colors.teal[200],
                direction: "2,1",
              ),
              ColoredWidget(
                color: Colors.teal[300],
                direction: "2,2",
              ),
              ColoredWidget(
                color: Colors.teal[400],
                direction: "2,3",
              ),
            ],
          ),
        ],
      ),
    );
  }
}

class PageControllerC {
  PageController controller;
  int recorded;
  PageControllerC({
    this.recorded,
    this.controller,
  });
}

class ColPageView extends StatefulWidget {
  final List<Widget> children;
  final List<PageControllerC> controllers;
  final ValueNotifier<int> notifier;
  final int idx;

  const ColPageView({
    Key key,
    this.children = const <Widget>[],
    @required this.idx,
    @required this.notifier,
    @required this.controllers,
  }) : super(key: key);

  @override
  _ColPageViewState createState() => _ColPageViewState();
}

class _ColPageViewState extends State<ColPageView> {
  @override
  Widget build(BuildContext context) {
    return PageView(
      controller: widget.controllers[widget.idx].controller,
      //   controller: widget.controller,
      scrollDirection: Axis.vertical,
      children: widget.children,
      onPageChanged: (widget.notifier.value == widget.idx)
          ? (pno) {
              var rand = Random();
              var randnn = rand.nextDouble();
              widget.controllers.forEach((colpv) {
                if (widget.controllers[widget.idx] == colpv) {
                  print("same widget so return $randnn");
                  return;
                }
                bool isSelected = colpv.controller.hasClients
                    ? colpv.controller.page == pno
                    : colpv.controller.initialPage == pno;

                if (!isSelected) {
                  print("not selected");
                  if (colpv.controller.hasClients) {
                    colpv.controller.animateToPage(pno,
                        duration: Duration(milliseconds: 200),
                        curve: Curves.easeIn);
                  }
                }
                print("$pno $isSelected");
              });
              print("col-${widget.idx} changed to $pno");
              widget.notifier.value = null;
            }
          : null,
    );
  }
}

class ColoredWidget extends StatefulWidget {
  final Color color;
  final String direction;

  const ColoredWidget({
    Key key,
    @required this.color,
    @required this.direction,
  }) : super(key: key);

  @override
  _ColoredWidgetState createState() => _ColoredWidgetState();
}

class _ColoredWidgetState extends State<ColoredWidget>
    with AutomaticKeepAliveClientMixin<ColoredWidget> {
  @override
  Widget build(BuildContext context) {
    super.build(context);
    return Container(
        color: widget.color,
        child: Center(
          child: Text(
            widget.direction,
            style: TextStyle(
              fontSize: 100,
              color: Colors.black,
            ),
          ),
        ));
  }

  @override
  bool get wantKeepAlive => true;
}
