import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';

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

  @override
  void initState() {
    _controllers = [
      PageController(keepPage: true),
      PageController(keepPage: true),
      PageController(keepPage: true),
      PageController(keepPage: true),
    ];
    _rowController = PageController(keepPage: true);
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: PageView(
        controller: _rowController,
        children: [
          ColPageView(
            controller: _controllers[0],
            idx: 0,
            children: <Widget>[
              ColoredWidget(
                color: Colors.cyan,
                direction: ">",
              ),
              ColoredWidget(
                color: Colors.orange,
                direction: ">>",
              ),
            ],
          ),
          ColPageView(
            idx: 1,
            controller: _controllers[1],
            children: [
              ColoredWidget(
                color: Colors.green,
                direction: "<",
              ),
              ColoredWidget(
                color: Colors.yellow,
                direction: "<<",
              ),
            ],
          ),
        ],
      ),
    );
  }
}

class ColPageView extends StatefulWidget {
  final List<Widget> children;
  final PageController controller;
  final int idx;

  const ColPageView({
    Key key,
    this.children = const <Widget>[],
    this.controller,
    @required this.idx,
  }) : super(key: key);

  @override
  _ColPageViewState createState() => _ColPageViewState();
}

class _ColPageViewState extends State<ColPageView> {
  @override
  Widget build(BuildContext context) {
    return PageView(
      controller: widget.controller,
      scrollDirection: Axis.vertical,
      children: widget.children,
      onPageChanged: (pno) {
        print("col-${widget.idx} changed to $pno");
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
