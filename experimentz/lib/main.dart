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
  List<PageController> _controllers;
  PageController _rowController;
  ValueNotifier<int> currIdxNotifier = ValueNotifier(0);
  ValueNotifier<int> currUpNotifier = ValueNotifier(0);

  @override
  void initState() {
    _controllers = [
      PageController(keepPage: true),
      PageController(keepPage: true),
      PageController(keepPage: true),
    ];
    _rowController = PageController(
      // initialPage: 0,
      // keepPage: true,
    );
		currIdxNotifier.value = _rowController.initialPage;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: PageView(
        pageSnapping: true,
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
            currup: currUpNotifier,
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
            currup: currUpNotifier,
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
            currup: currUpNotifier,
          ),
        ],
      ),
    );
  }
}

class ColPageView extends StatefulWidget {
  final List<Widget> children;
  final List<PageController> controllers;
  final ValueNotifier<int> notifier;
  final ValueNotifier<int> currup;
  final int idx;

  const ColPageView({
    Key key,
    this.children = const <Widget>[],
    @required this.idx,
    @required this.currup,
    @required this.notifier,
    @required this.controllers,
  }) : super(key: key);

  @override
  _ColPageViewState createState() => _ColPageViewState();
}

class _ColPageViewState extends State<ColPageView> {
  @override
  void initState() {
    print("INIT STATE ${widget.idx}");
    widget.controllers[widget.idx] =
        PageController(initialPage: widget.currup.value ?? 0);
    print("INIT STATE ${widget.currup.value}");
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return PageView(
      pageSnapping: true,
      controller: widget.controllers[widget.idx],
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
                bool isSelected = colpv.hasClients
                    ? colpv.page == pno
                    : colpv.initialPage == pno;

                if (!isSelected) {
                  print("not selected");
                  if (colpv.hasClients) {
                    colpv.animateToPage(pno,
                        duration: Duration(milliseconds: 200),
                        curve: Curves.easeIn);
                  }
                }
                widget.currup.value = pno;
                print("$pno $isSelected");
              });
              print("col-${widget.idx} changed to $pno");
              widget.notifier.value = null;
            }
          : (_) {
              print("nope ${widget.notifier.value} == ${widget.idx}");
            },
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
