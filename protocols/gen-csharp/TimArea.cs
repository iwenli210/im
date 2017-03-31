/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
using System;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using System.IO;
using Thrift;
using Thrift.Collections;
using System.Runtime.Serialization;
using Thrift.Protocol;
using Thrift.Transport;


#if !SILVERLIGHT
[Serializable]
#endif
public partial class TimArea : TBase
{
  private string _country;
  private string _province;
  private string _city;
  private List<TimNode> _extraList;
  private Dictionary<string, string> _extraMap;

  /// <summary>
  /// 国家
  /// </summary>
  public string Country
  {
    get
    {
      return _country;
    }
    set
    {
      __isset.country = true;
      this._country = value;
    }
  }

  /// <summary>
  /// 省
  /// </summary>
  public string Province
  {
    get
    {
      return _province;
    }
    set
    {
      __isset.province = true;
      this._province = value;
    }
  }

  /// <summary>
  /// 市
  /// </summary>
  public string City
  {
    get
    {
      return _city;
    }
    set
    {
      __isset.city = true;
      this._city = value;
    }
  }

  /// <summary>
  /// 附加信息：如区，镇，乡等
  /// </summary>
  public List<TimNode> ExtraList
  {
    get
    {
      return _extraList;
    }
    set
    {
      __isset.extraList = true;
      this._extraList = value;
    }
  }

  public Dictionary<string, string> ExtraMap
  {
    get
    {
      return _extraMap;
    }
    set
    {
      __isset.extraMap = true;
      this._extraMap = value;
    }
  }


  public Isset __isset;
  #if !SILVERLIGHT
  [Serializable]
  #endif
  public struct Isset {
    public bool country;
    public bool province;
    public bool city;
    public bool extraList;
    public bool extraMap;
  }

  public TimArea() {
  }

  public void Read (TProtocol iprot)
  {
    iprot.IncrementRecursionDepth();
    try
    {
      TField field;
      iprot.ReadStructBegin();
      while (true)
      {
        field = iprot.ReadFieldBegin();
        if (field.Type == TType.Stop) { 
          break;
        }
        switch (field.ID)
        {
          case 1:
            if (field.Type == TType.String) {
              Country = iprot.ReadString();
            } else { 
              TProtocolUtil.Skip(iprot, field.Type);
            }
            break;
          case 2:
            if (field.Type == TType.String) {
              Province = iprot.ReadString();
            } else { 
              TProtocolUtil.Skip(iprot, field.Type);
            }
            break;
          case 3:
            if (field.Type == TType.String) {
              City = iprot.ReadString();
            } else { 
              TProtocolUtil.Skip(iprot, field.Type);
            }
            break;
          case 4:
            if (field.Type == TType.List) {
              {
                ExtraList = new List<TimNode>();
                TList _list18 = iprot.ReadListBegin();
                for( int _i19 = 0; _i19 < _list18.Count; ++_i19)
                {
                  TimNode _elem20;
                  _elem20 = new TimNode();
                  _elem20.Read(iprot);
                  ExtraList.Add(_elem20);
                }
                iprot.ReadListEnd();
              }
            } else { 
              TProtocolUtil.Skip(iprot, field.Type);
            }
            break;
          case 5:
            if (field.Type == TType.Map) {
              {
                ExtraMap = new Dictionary<string, string>();
                TMap _map21 = iprot.ReadMapBegin();
                for( int _i22 = 0; _i22 < _map21.Count; ++_i22)
                {
                  string _key23;
                  string _val24;
                  _key23 = iprot.ReadString();
                  _val24 = iprot.ReadString();
                  ExtraMap[_key23] = _val24;
                }
                iprot.ReadMapEnd();
              }
            } else { 
              TProtocolUtil.Skip(iprot, field.Type);
            }
            break;
          default: 
            TProtocolUtil.Skip(iprot, field.Type);
            break;
        }
        iprot.ReadFieldEnd();
      }
      iprot.ReadStructEnd();
    }
    finally
    {
      iprot.DecrementRecursionDepth();
    }
  }

  public void Write(TProtocol oprot) {
    oprot.IncrementRecursionDepth();
    try
    {
      TStruct struc = new TStruct("TimArea");
      oprot.WriteStructBegin(struc);
      TField field = new TField();
      if (Country != null && __isset.country) {
        field.Name = "country";
        field.Type = TType.String;
        field.ID = 1;
        oprot.WriteFieldBegin(field);
        oprot.WriteString(Country);
        oprot.WriteFieldEnd();
      }
      if (Province != null && __isset.province) {
        field.Name = "province";
        field.Type = TType.String;
        field.ID = 2;
        oprot.WriteFieldBegin(field);
        oprot.WriteString(Province);
        oprot.WriteFieldEnd();
      }
      if (City != null && __isset.city) {
        field.Name = "city";
        field.Type = TType.String;
        field.ID = 3;
        oprot.WriteFieldBegin(field);
        oprot.WriteString(City);
        oprot.WriteFieldEnd();
      }
      if (ExtraList != null && __isset.extraList) {
        field.Name = "extraList";
        field.Type = TType.List;
        field.ID = 4;
        oprot.WriteFieldBegin(field);
        {
          oprot.WriteListBegin(new TList(TType.Struct, ExtraList.Count));
          foreach (TimNode _iter25 in ExtraList)
          {
            _iter25.Write(oprot);
          }
          oprot.WriteListEnd();
        }
        oprot.WriteFieldEnd();
      }
      if (ExtraMap != null && __isset.extraMap) {
        field.Name = "extraMap";
        field.Type = TType.Map;
        field.ID = 5;
        oprot.WriteFieldBegin(field);
        {
          oprot.WriteMapBegin(new TMap(TType.String, TType.String, ExtraMap.Count));
          foreach (string _iter26 in ExtraMap.Keys)
          {
            oprot.WriteString(_iter26);
            oprot.WriteString(ExtraMap[_iter26]);
          }
          oprot.WriteMapEnd();
        }
        oprot.WriteFieldEnd();
      }
      oprot.WriteFieldStop();
      oprot.WriteStructEnd();
    }
    finally
    {
      oprot.DecrementRecursionDepth();
    }
  }

  public override string ToString() {
    StringBuilder __sb = new StringBuilder("TimArea(");
    bool __first = true;
    if (Country != null && __isset.country) {
      if(!__first) { __sb.Append(", "); }
      __first = false;
      __sb.Append("Country: ");
      __sb.Append(Country);
    }
    if (Province != null && __isset.province) {
      if(!__first) { __sb.Append(", "); }
      __first = false;
      __sb.Append("Province: ");
      __sb.Append(Province);
    }
    if (City != null && __isset.city) {
      if(!__first) { __sb.Append(", "); }
      __first = false;
      __sb.Append("City: ");
      __sb.Append(City);
    }
    if (ExtraList != null && __isset.extraList) {
      if(!__first) { __sb.Append(", "); }
      __first = false;
      __sb.Append("ExtraList: ");
      __sb.Append(ExtraList);
    }
    if (ExtraMap != null && __isset.extraMap) {
      if(!__first) { __sb.Append(", "); }
      __first = false;
      __sb.Append("ExtraMap: ");
      __sb.Append(ExtraMap);
    }
    __sb.Append(")");
    return __sb.ToString();
  }

}

